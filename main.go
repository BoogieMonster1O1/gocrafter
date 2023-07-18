package main

import (
	"context"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
	_ "github.com/lib/pq"
	"gocrafter/handlers"
	"gocrafter/handlers/middleware"
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

	db, err := initDatabase()
	if err != nil {
		log.Fatal("Error initializing database", err)
	}
	store := session.New()
	app.Use("/app/**/*", middleware.SessionValidator(db, store))

	app.Get("/login", handlers.LoginHandler)
	app.Post("/login", handlers.LoginHandlerPost(db, store))
	app.Get("/logout", handlers.LogoutHandler(store))

	app.Get("/app", handlers.DashboardHandler)
	app.Get("/app/hosts", handlers.ManageHostsHandler(db, store))
	app.Delete("/app/hosts/delete/:id", handlers.ManageHostsDeleteHandler(db, store))

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
