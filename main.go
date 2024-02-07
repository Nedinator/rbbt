package main

import (
	"log"

	"github.com/Nedinator/ribbit/blogs"
	"github.com/Nedinator/ribbit/data"
	"github.com/Nedinator/ribbit/handlers"
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

	app.Post("/api/new-url", handlers.CreateURL)
	app.Get("/stats/:id", handlers.GetUrlStats)
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.Render("about", fiber.Map{})
	})
	app.Get("/blog", func(c *fiber.Ctx) error {
		posts, err := blogs.GetBlogPosts()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.Render("blog", fiber.Map{
			"Posts": posts,
		})
	})

	app.Get("/new-url", func(c *fiber.Ctx) error {
		return c.Render("new-url", fiber.Map{})
	})
	app.Get("/search", handlers.SearchForStats)

	app.Get("/:id", handlers.Redirect)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("home", fiber.Map{})
	})

	app.Listen(":3000")
}
