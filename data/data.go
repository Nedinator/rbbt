package data

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Url struct {
	ShortUrl  string    `json:"shorturl"`
	ShortId   string    `json:"shortid"`
	LongUrl   string    `json:"longurl"`
	Clicks    int       `json:"clicks"`
	CreatedAt time.Time `json:"createdat"`
	Owner     string    `json:"owner"`
	Referer   []Referer `json:"referer"`
}

type Referer struct {
	Domain string `json:"domain" bson:"domain"`
	Clicks int    `json:"clicks" bson:"clicks"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func AuthData(c *fiber.Ctx) fiber.Map {
	return fiber.Map{
		"IsLoggedIn": c.Locals("IsLoggedIn"),
		"Username":   c.Locals("Username"),
		"UserID":     c.Locals("id"),
	}
}
