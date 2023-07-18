package host

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gocrafter/models/data"
	"log"
	"net/http"
	"strconv"
)

func ManageHostsHandler(db *sql.DB, store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rows, err := db.Query("SELECT * FROM hosts")
		if err != nil {
			log.Println(err)
			return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
		}
		defer rows.Close()

		var hosts []data.Host

		for rows.Next() {
			var host data.Host
			err := rows.Scan(&host.Name, &host.ID, &host.SSHHostname, &host.SSHPort, &host.IsLocal, &host.Status)
			if err != nil {
				log.Println(err)
				return c.Redirect("/500.html")
			}

			hosts = append(hosts, host)
		}

		if err := rows.Err(); err != nil {
			log.Println(err)
			return c.Redirect("/500.html")
		}

		return c.Render("hosts", fiber.Map{
			"Page":  "/app/hosts",
			"Hosts": hosts,
		})
	}
}

func ManageHostsDeleteHandler(db *sql.DB, store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
		}

		var isLocal bool
		err := db.QueryRow("SELECT is_local FROM hosts WHERE id = $1", id).Scan(&isLocal)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.Status(fiber.StatusNotFound).SendString("Not Found")
			}
			log.Println(err.Error())
			return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
		}

		if isLocal {
			return c.Status(fiber.StatusForbidden).SendString("Cannot delete local host")
		}

		_, err = db.Exec("DELETE FROM hosts WHERE id = $1", id)
		if err != nil {
			log.Println(err.Error())
			return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
		}

		return c.SendStatus(fiber.StatusResetContent)
	}
}

func ManageHostsPostHandler(db *sql.DB, store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.FormValue("name", "No name")
		sshHostname := c.FormValue("sshHostname")
		sshPort := c.FormValue("sshPort")
		if sshHostname == "" || sshPort == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
		}
		sshPortNum, err := strconv.Atoi(sshPort)
		if err != nil {
			log.Println(err.Error())
			return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
		}

		_, err = db.Exec("INSERT INTO hosts (name, ssh_hostname, ssh_port, is_local) VALUES ($1, $2, $3, $4)", name, sshHostname, sshPortNum, false)
		if err != nil {
			log.Println(err.Error())
			return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
		}

		return c.SendStatus(fiber.StatusResetContent)
	}
}

func ManageHostsPatchHandler(db *sql.DB, store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		name := c.FormValue("name", "No name")
		if name == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
		}

		_, err := db.Exec("UPDATE hosts SET name = $1 WHERE id = $2", name, id)
		if err != nil {
			log.Println(err.Error())
			return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
		}

		return c.SendStatus(fiber.StatusResetContent)
	}
}
