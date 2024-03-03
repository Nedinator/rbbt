package data

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

const dbName = "rbbt"

var mongoURI = os.Getenv("MONGO_URI")
var Db *mongo.Database

func Connect() (*mongo.Client, *mongo.Database, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, nil, err
	}

	Db = client.Database(dbName)

	return client, Db, nil
}

func Disconnect(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatal(err)
	}
}
