package host

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gocrafter/lib"
	"log"
	"net/http"
)

func DirectoryListingHandler(db *sql.DB, store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		path := c.FormValue("path", "/")
		rows, err := db.Query("SELECT (is_local) FROM hosts WHERE hosts.id = $1", id)
		if err != nil {
			log.Println(err)
			return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
		}
		defer rows.Close()
		isLocal := false
		rows.Next()
		err = rows.Scan(&isLocal)
		if err != nil {
			log.Println(err)
			return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
		}

		if !isLocal {
			return c.Status(http.StatusNotImplemented).SendString("[]")
		}

		childItems, err := lib.GetChildItems(path)
		if err != nil {
			log.Println(err)
			return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
		}

		return c.JSON(childItems)
	}
}
