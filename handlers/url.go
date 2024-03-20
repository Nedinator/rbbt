package handlers

import (
	"log"
	"net/url"
	"os"
	"time"

	"github.com/Nedinator/ribbit/data"
	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

func GetUrlStats(c *fiber.Ctx) error {
	urlParams := c.Params("id")
	var res data.Url
	if err := data.DB().Where("short_id = ?", urlParams).First(&res).Error; err != nil {
		return c.Status(404).Render("404", data.AuthData(c))
	}
	statsData := data.AuthData(c)
	statsData["url"] = res
	return c.Render("stats", statsData)
}

func CreateURL(c *fiber.Ctx) error {
	newurl := data.Url{
		LongUrl: c.FormValue("longurl"),
		Owner:   c.Locals("Username").(string),
		ShortId: c.FormValue("shortid"),
		Clicks:  0,
	}

	if newurl.ShortId == "" {
		generatedId, err := gonanoid.New(6)
		if err != nil {
			return c.Status(500).SendString("Internal Server Error.")
		} else {
			newurl.ShortId = generatedId
		}
	}
	newurl.ShortUrl = os.Getenv("DOMAIN") + "/" + newurl.ShortId

	if err := data.DB().Create(&newurl).Error; err != nil {
		return c.Status(500).SendString("Internal Server Error.")
	}

	nextPageData := data.AuthData(c)
	nextPageData["url"] = newurl
	return c.Render("stats", nextPageData)
}

func Redirect(c *fiber.Ctx) error {
	var res data.Url
	urlParams := c.Params("id")

	err := data.DB().Model(&data.Url{}).Where("short_id = ?", urlParams).UpdateColumn("clicks", gorm.Expr("clicks + ?", 1)).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).Render("404", data.AuthData(c))
		}
		return c.Status(500).SendString("Internal Server Error")
	}

	err = data.DB().Where("short_id = ?", urlParams).First(&res).Error
	if err != nil {
		log.Panic(err)
		return c.Status(500).SendString("Internal Server Error")
	}

	if res.Referers == nil {
		res.Referers = make(map[string]data.Referer)
	}

	rawURL := c.Get("Referer")
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		log.Println("Error parsing referer URL:", err)
	} else {
		domain := parsedURL.Hostname()
		if domain != "" {
			referer, exists := res.Referers[domain]
			if exists {
				referer.Clicks++
				res.Referers[domain] = referer
			} else {
				newReferer := data.Referer{Domain: domain, Clicks: 1, Tags: make(map[string]data.Tag)}
				res.Referers[domain] = newReferer
			}
			data.DB().Save(&res)
		}
	}

	return c.Redirect(res.LongUrl)
}

func SearchForStats(c *fiber.Ctx) error {
	searchID := c.Query("searchid")
	var res data.Url
	if err := data.DB().Where("short_id = ?", searchID).First(&res).Error; err != nil {
		return c.Status(404).Render("404", data.AuthData(c))
	}
	timzoneCookie := c.Cookies("timezone", "UTC")
	loc, err := time.LoadLocation(timzoneCookie)
	if err != nil {
		loc = time.UTC
	}
	res.CreatedAt = convertToLocalTimeUsingLocation(res.CreatedAt, loc)
	nextPageData := data.AuthData(c)
	nextPageData["url"] = res
	return c.Render("stats", nextPageData)
}

func DeleteUrl(c *fiber.Ctx) error {
	urlParams := c.Params("id")
	if err := data.DB().Where("short_id = ?", urlParams).Delete(&data.Url{}).Error; err != nil {
		return c.Status(500).SendString("Failed to delete URL. Give that another go...")
	}
	return c.Redirect("/dashboard")
}

func convertToLocalTimeUsingLocation(t time.Time, loc *time.Location) time.Time {
	return t.In(loc)
}
