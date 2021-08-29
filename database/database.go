package database

import (
	"context"
	"kardashian_api/config"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var DB *mongo.Database

func connectClient() error {
	ctx, cancel := Context()
	defer cancel()

	err := client.Connect(ctx)
	if err != nil {
		return err
	}
	log.Printf("✅ MongoDB Client: connect success")
	return nil
}

func newClient() error {
	c, err := mongo.NewClient(options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		return err
	}
	client = c
	log.Printf("✅ MongoDB Client: client success")
	return nil
}

func dbSetup() {
	DB = client.Database(config.MongoDBNAME)

}
func LoadClient() error {
	var err error
	if client == nil {
		err = newClient()
		if err != nil {
			return err
		}
		err = connectClient()
		if err != nil {
			return err
		}
	}
	dbSetup()
	return nil
}

func Context() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 3*time.Second)
}

func ListOfCollections() ([]string, error) {
	ctx, cancel := Context()
	defer cancel()

	list, err := DB.ListCollectionNames(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	return list, err
}

func CountCollectionDocs(c string, filter bson.M) int {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	opts := options.Count().SetMaxTime(2 * time.Second)
	total, err := DB.Collection(c).CountDocuments(ctx, filter, opts)
	if err != nil {
		log.Printf("Error CountCollectionDocs %v - %v", c, err)
	}
	return int(total)
}
