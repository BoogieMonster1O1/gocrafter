package middleware

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"log"
)

func SessionValidator(db *sql.DB, store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			panic("noooooo")
		}
		id := sess.Get("id")
		if id == nil {
			return c.Redirect("/401.html")
		}

		var count int
		err = db.QueryRow("SELECT COUNT(*) FROM users WHERE id = $1", id).Scan(&count)
		if err != nil {
			log.Println("Error querying the database:", err)
			return c.Redirect("/500.html")
		} else if count == 0 {
			return c.Redirect("/401.html")
		}

		return c.Next()
	}
}
