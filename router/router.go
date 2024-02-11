// TODO: Once I finish more systems like auth, I will move all the routes here.
package router

import (
	"github.com/Nedinator/ribbit/blogs"
	"github.com/Nedinator/ribbit/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/api/new-url", handlers.CreateURL)
	app.Post("/auth/signup", handlers.Signup)
	app.Post("/auth/login", handlers.Login)
	app.Get("/stats/:id", handlers.GetUrlStats)
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.Render("about", fiber.Map{})
	})
	app.Get("/blog", func(c *fiber.Ctx) error {
		posts, err := blogs.GetBlogPosts()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).Render("blog", fiber.Map{})
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
}
