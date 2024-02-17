package dashboard

import (
	"context"
	"time"

	"github.com/Nedinator/ribbit/data"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetLinks(c *fiber.Ctx) ([]data.Url, error) {
	var links []data.Url
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	owner := c.Locals("username").(string)
	if owner == "" {
		return nil, nil
	}

	filter := bson.M{"owner": owner}
	findOptions := options.Find().SetSort(bson.D{{"createdat", -1}}) // Sort by creation date, newest first

	cursor, err := data.Db.Collection("url").Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var link data.Url
		if err = cursor.Decode(&link); err != nil {
			return nil, err
		}
		links = append(links, link)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return links, nil
}
