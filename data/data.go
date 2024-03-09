package data

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	ShortUrl string `gorm:"column:short_url"`
	ShortId  string `gorm:"column:short_id"`
	LongUrl  string `gorm:"column:long_url"`
	Clicks   int
	Owner    string
	Referers map[string]Referer `gorm:"-"`
}

type Referer struct {
	UrlID  uint `gorm:"column:url_id;foreignKey:url_id"`
	Domain string
	Clicks int
	Tags   map[string]Tag `gorm:"-"`
}

type Tag struct {
	Name  string
	Color string
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
