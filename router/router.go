package router

import (
	"github.com/Nedinator/ribbit/data"
	"github.com/Nedinator/ribbit/handlers"
	"github.com/Nedinator/ribbit/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Use(middleware.AuthStatusMiddleware)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("home", data.AuthData(c))
	})
	app.Post("/api/new-url", handlers.CreateURL)
	app.Post("/auth/signup", handlers.Signup)
	app.Post("/auth/login", handlers.Login)
	app.Post("/auth/logout", handlers.Logout)
	app.Get("/stats/:id", handlers.GetUrlStats)
	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", data.AuthData(c))
	})
	app.Get("/signup", func(c *fiber.Ctx) error {
		return c.Render("signup", data.AuthData(c))
	})
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.Render("about", data.AuthData(c))
	})

	app.Get("/dashboard", middleware.JwtMiddleware, func(c *fiber.Ctx) error {
		return c.Render("dashboard", data.AuthData(c))
	})

	app.Get("/new-url", func(c *fiber.Ctx) error {
		return c.Render("new-url", data.AuthData(c))

	})

	app.Get("/search", handlers.SearchForStats)
	app.Get("/:id", handlers.Redirect)

}
