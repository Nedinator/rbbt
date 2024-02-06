package handlers

import (
	"net/url"
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

	return c.Render("stats", res)
}

func CreateURL(c *fiber.Ctx) error {
	var newurl data.Url

	longurl := c.FormValue("longurl")
	newurl.LongUrl = longurl
	if !isValidURL(newurl.LongUrl) {
		return c.Status(400).SendString("Invalid URL")
	}

	shortid, err := gonanoid.New(6)
	if err != nil {
		return c.Status(500).SendString("Internal Server Error")
	}
	newurl.ShortId = shortid
	newurl.ShortUrl = "http://127.0.0.1:3000/" + shortid
	newurl.Clicks = 0
	newurl.CreatedAt = time.Now()
	data.Db.Collection("url").InsertOne(c.Context(), newurl)
	return c.Render("stats", newurl)
}

func Redirect(c *fiber.Ctx) error {
	var res data.Url
	urlParams := c.Params("id")
	filter := bson.M{"shortid": urlParams}
	update := bson.M{"$inc": bson.M{"clicks": 1}}

	err := data.Db.Collection("url").FindOneAndUpdate(c.Context(), filter, update).Decode(&res)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(404).SendString("Not Found")
		}
		return c.Status(500).SendString("Internal Server Error")
	}

	c.Redirect(res.LongUrl)
	return nil
}

func SearchForStats(c *fiber.Ctx) error {
	var res data.Url
	searchID := c.FormValue("searchid")
	filter := bson.M{"shortid": searchID}
	err := data.Db.Collection("url").FindOne(c.Context(), filter).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(404).Render("404", nil)
		}
		return c.Status(500).Render("404", nil) // Todo: change to 500 page
	}

	return c.Render("stats", res)
}

func isValidURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	return err == nil
}
