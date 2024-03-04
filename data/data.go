package data

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	ShortUrl  string    `json:"shorturl" gorm:"column:shorturl;unique;not null"`
	ShortId   string    `json:"shortid" gorm:"column:shortid;unique;not null"`
	LongUrl   string    `json:"longurl" gorm:"column:longurl;not null"`
	Clicks    int       `json:"clicks" gorm:"column:clicks;default:0"`
	CreatedAt time.Time `json:"createdat" gorm:"column:createdat"`
	Owner     string    `json:"owner" gorm:"column:owner"`
	Referer   []Referer `json:"referers" gorm:"column:referers;"`
}

type Referer struct {
	Domain string   `json:"domain"`
	Clicks int      `json:"clicks"`
	Tags   []string `json:"tags"`
}

type User struct {
	gorm.Model
	ID       string `json:"id"`
	Username string `json:"username" gorm:"column:username;unique;not null"`
	Password string `json:"password" gorm:"column:password;not null"`
}

func AuthData(c *fiber.Ctx) fiber.Map {
	return fiber.Map{
		"IsLoggedIn": c.Locals("IsLoggedIn"),
		"Username":   c.Locals("Username"),
		"UserID":     c.Locals("id"),
	}
}
