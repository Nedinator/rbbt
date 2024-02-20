package data

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Url struct {
	ShortUrl  string    `json:"shorturl" bson:"shorturl"`
	ShortId   string    `json:"shortid" bson:"shortid"`
	LongUrl   string    `json:"longurl" bson:"longurl"`
	Clicks    int       `json:"clicks" bson:"clicks"`
	CreatedAt time.Time `json:"createdat" bson:"createdat"`
	Owner     string    `json:"owner" bson:"owner"`
	Referer   []Referer `json:"referer" bson:"referer"`
}

type Referer struct {
	Domain string   `json:"domain" bson:"domain"`
	Clicks int      `json:"clicks" bson:"clicks"`
	Tags   []string `json:"tags" bson:"tags"`
}

type User struct {
	ID       string `json:"id" bson:"id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

func AuthData(c *fiber.Ctx) fiber.Map {
	return fiber.Map{
		"IsLoggedIn": c.Locals("IsLoggedIn"),
		"Username":   c.Locals("Username"),
		"UserID":     c.Locals("id"),
	}
}
