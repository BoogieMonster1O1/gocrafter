package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gocrafter/models"
)

func LoginHandler(c *fiber.Ctx) error {
	expired := c.Query("expired")

	loginForm := models.LoginForm{
		Expired: expired,
		Invalid: "",
	}

	return c.Render("login", loginForm)
}

func LoginHandlerPost(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	authSuccess := authenticateUser(email, password)

	if !authSuccess {

		loginForm := models.LoginForm{
			Expired: "",
			Invalid: "true",
		}

		return c.Render("login", loginForm)
	}

	return c.Redirect("/app")
}

// Placeholder function to authenticate the user
func authenticateUser(email, password string) bool {
	// Perform authentication logic here
	// Replace with your actual authentication code
	// Return true if authentication is successful, false otherwise
	return false
}
