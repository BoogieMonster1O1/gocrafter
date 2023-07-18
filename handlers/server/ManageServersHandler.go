package server

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gocrafter/models"
	"log"
	"net/http"
)

func ManageServersHandler(db *sql.DB, store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rows, err := db.Query("SELECT s.id, s.absolute_path, s.name, h.name AS host_name FROM servers AS s JOIN hosts AS h ON s.host_id = h.id ORDER BY h.name, s.name;")
		if err != nil {
			log.Println(err)
			return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
		}
		defer rows.Close()

		var servers []models.ServerModel

		for rows.Next() {
			var server models.ServerModel
			err := rows.Scan(&server.Id, &server.AbsolutePath, &server.Name, &server.HostName)
			if err != nil {
				log.Println(err)
				return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
			}

			servers = append(servers, server)
		}

		if err := rows.Err(); err != nil {
			log.Println(err)
			return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
		}

		return c.Render("servers", fiber.Map{
			"Page":    "/app/servers",
			"Servers": servers,
		})
	}
}
