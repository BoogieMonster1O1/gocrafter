package server

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gocrafter/models"
	"gocrafter/models/data"
	"log"
	"net/http"
)

func CreateServerHandler(db *sql.DB, store *session.Store) fiber.Handler {
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

		return c.Render("create_server", fiber.Map{
			"Page":     "",
			"Hosts":    hosts,
			"Manifest": models.CachedVersionManifest,
		})
	}
}

func CreateServerPostHandler(db *sql.DB, store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}
