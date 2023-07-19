package main

import (
	"context"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
	_ "github.com/lib/pq"
	"gocrafter/handlers"
	"gocrafter/handlers/host"
	"gocrafter/handlers/keypair"
	"gocrafter/handlers/middleware"
	"gocrafter/handlers/onboarding"
	"gocrafter/handlers/server"
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

	app.Get("/login", onboarding.LoginHandler)
	app.Post("/login", onboarding.LoginHandlerPost(db, store))
	app.Get("/logout", onboarding.LogoutHandler(store))

	app.Get("/app", handlers.DashboardHandler)
	app.Get("/app/ssh-key", keypair.ViewSSHKeyHandler(db, store))
	app.Delete("/app/ssh-key/invalidate", keypair.InvalidateSSHKeyHandler(db, store))
	app.Get("/app/hosts", host.ManageHostsHandler(db, store))
	app.Post("/app/hosts/create", host.ManageHostsPostHandler(db, store))
	app.Delete("/app/hosts/:id/delete", host.ManageHostsDeleteHandler(db, store))
	app.Patch("/app/hosts/:id/edit", host.ManageHostsPatchHandler(db, store))
	app.Get("/app/hosts/:id/directory", host.DirectoryListingHandler(db, store))
	app.Get("/app/servers", server.ManageServersHandler(db, store))
	app.Get("/app/servers/create", server.CreateServerHandler(db, store))
	app.Post("/app/servers/create", server.CreateServerPostHandler(db, store))
	app.Get("/app/servers/import", server.ImportServerHandler(db, store))
	app.Get("/app/servers/:id", server.ServerDashboardHandler(db, store))

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
