package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func LogoutHandler(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return err
		}
		err = sess.Destroy()
		if err != nil {
			return err
		}
		return c.Redirect("/login")
	}
}
