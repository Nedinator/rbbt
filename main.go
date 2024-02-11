package main

import (
	"log"

	"github.com/Nedinator/ribbit/data"

	"github.com/Nedinator/ribbit/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	client, _, err := data.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer data.Disconnect(client)
	engine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	router.SetupRoutes(app)

	app.Listen(":3000")
}
