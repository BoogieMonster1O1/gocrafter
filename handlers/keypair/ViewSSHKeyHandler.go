package keypair

import (
	"database/sql"
	"encoding/base64"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gocrafter/lib"
	"log"
	"net/http"
)

func ViewSSHKeyHandler(db *sql.DB, store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		key, err := lib.GetPublicKey()
		if err != nil {
			log.Println(err.Error())
			return c.Status(http.StatusInternalServerError).SendString("Internal server error")
		}
		encodedKey := base64.StdEncoding.EncodeToString(key)
		return c.Render("sshkey", fiber.Map{
			"PublicKey": encodedKey,
		})
	}
}
