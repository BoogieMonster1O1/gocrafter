package server

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func CreateServerHandler(db *sql.DB, store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("no")
	}
}
