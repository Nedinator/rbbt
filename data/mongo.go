package data

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const dbName = "rbbt"
const mongoURI = "mongodb://localhost:27017/" + dbName

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
