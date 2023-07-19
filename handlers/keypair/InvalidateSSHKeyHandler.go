package keypair

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gocrafter/lib"
	"log"
	"net/http"
)

func InvalidateSSHKeyHandler(db *sql.DB, store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := lib.GenerateKeyPair()
		if err != nil {
			log.Println(err.Error())
			return c.Status(http.StatusInternalServerError).SendString("Internal server error")
		}

		return c.SendStatus(http.StatusNoContent)
	}
}
