package main

import (
	"os"

	"github.com/Nedinator/ribbit/data"

	"github.com/Nedinator/ribbit/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

	dsn := os.Getenv("DB_DSN")
	data.OpenDB(dsn)

	engine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	router.SetupRoutes(app)

	app.Listen(":3000")
}
