package handlers

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gocrafter/models/data"
	"log"
	"net/http"
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
			err := rows.Scan(&host.Name, &host.ID, &host.SSHHostname, &host.SSHPort, &host.IsLocal)
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
