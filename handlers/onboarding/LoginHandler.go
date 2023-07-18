package onboarding

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gocrafter/models"
	"gocrafter/models/data"
	"log"
)

func LoginHandler(c *fiber.Ctx) error {
	expired := c.Query("expired")

	loginForm := models.LoginForm{
		Expired: expired,
		Invalid: "",
	}

	return c.Render("login", loginForm)
}

func LoginHandlerPost(db *sql.DB, store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		email := c.FormValue("email")
		password := c.FormValue("password")
		success, user := authenticateUser(db, email, password)

		if !success {

			loginForm := models.LoginForm{
				Expired: "",
				Invalid: "true",
			}

			return c.Render("login", loginForm)
		}

		sess, err := store.Get(c)
		if err != nil {
			return c.SendStatus(503)
		}

		sess.Set("id", user.ID)

		err = sess.Save()
		if err != nil {
			return c.SendStatus(503)
		}

		return c.Redirect("/app")
	}
}

func authenticateUser(db *sql.DB, email, password string) (bool, data.User) {
	var user data.User
	rows, err := db.Query("SELECT * FROM users WHERE users.email = $1", email)
	if err != nil {
		log.Print(err.Error())
		return false, user
	}
	log.Print("conc")
	rows.Next()
	if err := rows.Scan(&user.DisplayName, &user.Email, &user.Password, &user.Role, &user.ID); err != nil {
		log.Print(err.Error())
		return false, user
	}
	log.Print(user.Password)
	if err := rows.Close(); err != nil {
		log.Print(err.Error())
		return false, user
	}
	return user.Password == password, user
}
