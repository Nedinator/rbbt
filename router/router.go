package router

import (
	"github.com/Nedinator/ribbit/blogs"
	// "github.com/Nedinator/ribbit/dashboard"
	"github.com/Nedinator/ribbit/handlers"
	"github.com/Nedinator/ribbit/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Use(middleware.AuthStatusMiddleware)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("home", commonData(c))
	})
	app.Post("/api/new-url", handlers.CreateURL)
	app.Post("/auth/signup", handlers.Signup)
	app.Post("/auth/login", handlers.Login)
	app.Post("/auth/logout", handlers.Logout)
	app.Get("/stats/:id", handlers.GetUrlStats)
	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", commonData(c))
	})
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.Render("about", commonData(c))
	})

	app.Get("/blog", func(c *fiber.Ctx) error {
		posts, err := blogs.GetBlogPosts()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).Render("blog", fiber.Map{})
		}

		data := commonData(c)
		data["Posts"] = posts

		return c.Render("blog", data)
	})

	app.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.Render("dashboard", commonData(c))
	})

	app.Get("/new-url", func(c *fiber.Ctx) error {
		return c.Render("new-url", commonData(c))

	})

	app.Get("/search", handlers.SearchForStats)
	app.Get("/:id", handlers.Redirect)

}

func commonData(c *fiber.Ctx) fiber.Map {
	return fiber.Map{
		"IsLoggedIn": c.Locals("IsLoggedIn"),
		"Username":   c.Locals("Username"),
		"UserID":     c.Locals("id"),
	}
}
