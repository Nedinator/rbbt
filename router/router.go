package router

import (
	"fmt"

	"github.com/Nedinator/ribbit/blogs"
	"github.com/Nedinator/ribbit/dashboard"
	"github.com/Nedinator/ribbit/handlers"
	"github.com/Nedinator/ribbit/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Use(middleware.AuthStatusMiddleware)
	app.Get("/", func(c *fiber.Ctx) error {
		isLoggedIn, ok := c.Locals("IsLoggedIn").(bool)
		if !ok {
			isLoggedIn = false
		}

		var username string
		if isLoggedIn {
			usernameInterface, ok := c.Locals("username").(string)
			if !ok {
				fmt.Println("username not found in locals")
				username = ""
			} else {
				username = usernameInterface
			}
		}
		fmt.Println("username" + username)
		return c.Render("home", fiber.Map{
			"IsLoggedIn": isLoggedIn,
			"Username":   username,
		})
	})
	app.Post("/api/new-url", handlers.CreateURL)
	app.Post("/auth/signup", handlers.Signup)
	app.Post("/auth/login", handlers.Login)
	app.Post("/auth/logout", handlers.Logout)
	app.Get("/stats/:id", handlers.GetUrlStats)
	app.Get("/login", func(c *fiber.Ctx) error {
		isLoggedIn, ok := c.Locals("IsLoggedIn").(bool)
		if !ok {
			isLoggedIn = false
		}

		var username string
		if isLoggedIn {
			usernameInterface, ok := c.Locals("username").(string)
			if !ok {
				username = ""
			} else {
				username = usernameInterface
			}
		}
		return c.Render("login", fiber.Map{
			"IsLoggedIn": isLoggedIn,
			"Username":   username,
		})
	})
	app.Get("/about", func(c *fiber.Ctx) error {
		isLoggedIn, ok := c.Locals("IsLoggedIn").(bool)
		if !ok {
			isLoggedIn = false
		}

		var username string
		if isLoggedIn {
			usernameInterface, ok := c.Locals("username").(string)
			if !ok {
				username = ""
			} else {
				username = usernameInterface
			}
		}
		return c.Render("about", fiber.Map{
			"IsLoggedIn": isLoggedIn,
			"Username":   username,
		})
	})

	app.Get("/blog", func(c *fiber.Ctx) error {
		posts, err := blogs.GetBlogPosts()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).Render("blog", fiber.Map{})
		}
		isLoggedIn, ok := c.Locals("IsLoggedIn").(bool)
		if !ok {
			isLoggedIn = false
		}

		var username string
		if isLoggedIn {
			usernameInterface, ok := c.Locals("username").(string)
			if !ok {
				username = ""
			} else {
				username = usernameInterface
			}
		}
		return c.Render("blog", fiber.Map{
			"IsLoggedIn": isLoggedIn,
			"Username":   username,
			"Posts":      posts,
		})
	})

	app.Get("/new-url", func(c *fiber.Ctx) error {
		userIDInterface := c.Locals("userID")
		userID, ok := userIDInterface.(string)
		if !ok {
			userID = ""
		}
		isLoggedIn, ok := c.Locals("IsLoggedIn").(bool)
		if !ok {
			isLoggedIn = false
		}

		var username string
		if isLoggedIn {
			usernameInterface, ok := c.Locals("username").(string)
			if !ok {
				username = ""
			} else {
				username = usernameInterface
			}
		}
		return c.Render("new-url", fiber.Map{
			"IsLoggedIn": isLoggedIn,
			"Username":   username,
			"ID":         userID,
		})

	})
	app.Get("/search", handlers.SearchForStats)
	app.Get("/:id", handlers.Redirect)

	app.Get("/dashboard", func(c *fiber.Ctx) error {
		userIDInterface := c.Locals("userID")
		userID, ok := userIDInterface.(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}
		links, err := dashboard.GetLinks(userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
		isLoggedIn, ok := c.Locals("IsLoggedIn").(bool)
		if !ok {
			isLoggedIn = false
		}

		usernameInterface := c.Locals("username")
		username, ok := usernameInterface.(string)
		if !ok {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
		return c.Render("dashboard", fiber.Map{
			"Links":      links,
			"Username":   username,
			"IsLoggedIn": isLoggedIn,
		})
	})
}
