package main

import (
	"context"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	_ "github.com/lib/pq"
	"gocrafter/handlers"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	engine := html.NewFileSystem(http.Dir("./views"), ".html")
	engine.Reload(true)
	engine.Debug(true)
	engine.Layout("embed")
	engine.Delims("{{", "}}")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/static", "./static")

	dbPool, err := initDatabase()
	if err != nil {
		log.Fatal("Error initializing database", err)
	}
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", dbPool)
		return c.Next()
	})

	app.Get("/login", handlers.LoginHandler)
	app.Post("/login", handlers.LoginHandlerPost)

	log.Fatal(app.Listen(":3000"))
}

func initDatabase() (*sql.DB, error) {
	str := os.Getenv("DB_CONNECTION_STRING")
	db, err := sql.Open("postgres", str)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
