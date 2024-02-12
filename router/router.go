package router

import (
	"github.com/Nedinator/ribbit/blogs"
	"github.com/Nedinator/ribbit/dashboard"
	"github.com/Nedinator/ribbit/handlers"
	"github.com/Nedinator/ribbit/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("home", fiber.Map{})
	})
	app.Post("/api/new-url", handlers.CreateURL)
	app.Post("/auth/signup", handlers.Signup)
	app.Post("/auth/login", handlers.Login)
	app.Get("/stats/:id", handlers.GetUrlStats)
	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", fiber.Map{})
	})
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
	app.Get("/dashboard", middleware.JwtMiddleware, func(c *fiber.Ctx) error {
		userIDInterface := c.Locals("userID")
		userID, ok := userIDInterface.(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}
		links, err := dashboard.GetLinks(userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
		usernameInterface := c.Locals("username")
		username, ok := usernameInterface.(string)
		if !ok {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
		return c.Render("dashboard", fiber.Map{
			"Links":    links,
			"Username": username,
			"userID":   userID,
		})
	})
}
