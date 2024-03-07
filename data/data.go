package data

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	ShortUrl string
	ShortId  string
	LongUrl  string
	Clicks   int
	Owner    string
	Referers []Referer
}

type Referer struct {
	UrlID  uint `gorm:"foreignKey:urlid"`
	Domain string
	Clicks int
	Tags   []string
}

type User struct {
	gorm.Model
	Username string
	Password string
}

func AuthData(c *fiber.Ctx) fiber.Map {
	return fiber.Map{
		"IsLoggedIn": c.Locals("IsLoggedIn"),
		"Username":   c.Locals("Username"),
		"UserID":     c.Locals("id"),
	}
}
