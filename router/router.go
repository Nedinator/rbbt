package router

import (
	"errors"

	"github.com/Nedinator/ribbit/dashboard"
	"github.com/Nedinator/ribbit/data"
	"github.com/Nedinator/ribbit/handlers"
	"github.com/Nedinator/ribbit/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App) {
	app.Use(middleware.AuthStatusMiddleware)
	app.Use(logger.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("home", data.AuthData(c))
	})
	app.Post("/api/new-url", handlers.CreateURL)
	app.Post("/auth/signup", handlers.Signup)
	app.Post("/auth/login", handlers.Login)
	app.Post("/auth/logout", handlers.Logout)
	app.Get("/stats/:id", handlers.GetUrlStats)
	app.Get("/delete/:id", middleware.JwtMiddleware, handlers.DeleteUrl)
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
		urls, err := dashboard.GetLinks(c)
		if err != nil {
			return c.Status(500).SendString("Internal Server Error. If you see this you should prolly dial 911...")
		}
		data := data.AuthData(c)
		data["Links"] = urls
		return c.Render("dashboard", data)
	})

	app.Get("/dashboard/:id", middleware.JwtMiddleware, func(c *fiber.Ctx) error {
		var url data.Url
		renderData := data.AuthData(c)

		urlParams := c.Params("id")

		result := data.DB().Where("short_id = ?", urlParams).First(&url)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return c.Status(fiber.StatusNotFound).Render("404", nil) // Assuming you have a 404 template
			}

			return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{"message": "Internal Server Error"})
		}

		renderData["url"] = url
		return c.Render("stats", renderData) // Ensure you have a "stats" template
	})

	app.Get("/new-url", func(c *fiber.Ctx) error {
		return c.Render("new-url", data.AuthData(c))
	})

	app.Get("/search", handlers.SearchForStats)
	app.Get("/:id", handlers.Redirect)

}
