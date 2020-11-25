package storage

import (
	"context"
	"log"
	"os"

	"github.com/pjonathas/url-shortener/shortener"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()
var collection *mongo.Collection
var database *mongo.Database

func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_CONN_STR")))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database("shortener")
	collection = database.Collection("urls")
}

type Mongo struct{}

func (s Mongo) Put(shortenerURL shortener.URL) error {

	_, err := collection.InsertOne(ctx, shortenerURL)
	if err != nil {
		return err
	}

	return nil
}

func (s Mongo) Get(id string) (shortener.URL, error) {
	var shortenerURL shortener.URL

	collection.FindOne(ctx, bson.M{"_id": id}).Decode(&shortenerURL)

	return shortenerURL, nil
}
