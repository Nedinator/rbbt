package handlers

import (
	"net/url"
	"os"
	"time"

	"github.com/Nedinator/ribbit/data"
	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUrlStats(c *fiber.Ctx) error {
	urlParams := c.Params("id")
	filter := bson.M{"shortid": urlParams}
	var res data.Url
	err := data.Db.Collection("url").FindOne(c.Context(), filter).Decode(&res)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(404).Render("404", nil)
		}
		return c.Status(500).Render("404", nil)
	}
	statsData := data.AuthData(c)
	statsData["url"] = res
	return c.Render("stats", statsData)
}

func CreateURL(c *fiber.Ctx) error {
	var newurl data.Url

	longurl := c.FormValue("longurl")
	newurl.LongUrl = longurl
	username := c.Locals("Username").(string)
	newurl.Owner = username
	customId := c.FormValue("shortId")
	var shortid string
	if customId == "" {
		generatedId, err := gonanoid.New(6)
		if err != nil {
			return c.Status(500).SendString("Internal Server Error. If you see this you should prolly dial 911...")
		}
		shortid = generatedId
	} else {
		shortid = customId
	}
	newurl.ShortId = shortid
	newurl.ShortUrl = os.Getenv("DOMAIN") + "/" + shortid
	newurl.Clicks = 0
	newurl.Referer = []data.Referer{}
	newurl.CreatedAt = time.Now()
	data.Db.Collection("url").InsertOne(c.Context(), newurl)

	nextPageData := data.AuthData(c)
	nextPageData["url"] = newurl
	return c.Render("stats", nextPageData)
}

func Redirect(c *fiber.Ctx) error {
	var res data.Url
	urlParams := c.Params("id")
	filter := bson.M{"shortid": urlParams}
	update := bson.M{"$inc": bson.M{"clicks": 1}}

	err := data.Db.Collection("url").FindOneAndUpdate(c.Context(), filter, update).Decode(&res)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(404).Render("404", nil)
		}
		return c.Status(500).Render("404", nil)
	}

	rawURL := c.Get("Referer")
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return c.Status(500).SendString("Internal Server Error. If you see this you should prolly dial 911...")
	}
	domain := parsedURL.Hostname()
	if domain != "" {
		domainExists := false
		for _, ref := range res.Referer {
			if ref.Domain == domain {
				domainFilter := bson.M{"shortid": urlParams, "referer.domain": domain}
				domainUpdate := bson.M{"$inc": bson.M{"referer.$.clicks": 1}}
				_, err = data.Db.Collection("url").UpdateOne(c.Context(), domainFilter, domainUpdate)
				domainExists = true
				break
			}
		}

		if !domainExists {
			newReferer := bson.M{"$push": bson.M{"referer": data.Referer{Domain: domain, Clicks: 1}}}
			_, err = data.Db.Collection("url").UpdateOne(c.Context(), filter, newReferer)
		}
	}

	return c.Redirect(res.LongUrl)
}

func SearchForStats(c *fiber.Ctx) error {
	var res data.Url
	searchID := c.Query("searchid")
	filter := bson.M{"shortid": searchID}
	err := data.Db.Collection("url").FindOne(c.Context(), filter).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(404).Render("404", nil)
		}

		return c.Status(500).Render("404", nil)
	}
	timzoneCookie := c.Cookies("timezone", "UTC")
	loc, err := time.LoadLocation(timzoneCookie)
	if err != nil {
		loc = time.UTC
	}
	userCreatedAt, err := convertToLocalTimeUsingLocation(res.CreatedAt, loc)
	res.CreatedAt = userCreatedAt
	nextPageData := data.AuthData(c)
	nextPageData["url"] = res
	return c.Render("stats", nextPageData)
}

func convertToLocalTimeUsingLocation(t time.Time, loc *time.Location) (time.Time, error) {
	return t.In(loc), nil
}
