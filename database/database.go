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

func LoadClient() error {

	if client == nil {

		c, err := mongo.NewClient(options.Client().ApplyURI(config.MongoURI))
		if err != nil {
			return err
		}
		client = c
		log.Printf("MongoDB Client: success creating the client")

		ctx, cancel := Context()
		defer cancel()

		errConnect := client.Connect(ctx)
		if errConnect != nil {
			return errConnect
		}

	}

	DB = client.Database(config.MongoDBNAME)
	log.Printf("MongoDB Client: success connecting to the db")

	return nil
}

func Use(tableName string) *mongo.Collection {
	coll := DB.Collection(tableName)
	return coll
}

func Context() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

func ValidCollection(coll string) bool {
	collLst, errColl := Collections()

	if errColl != nil {
		return false
	}
	for _, c := range collLst {
		if c == coll {
			return true
		}
	}
	return false
}

func Collections() ([]string, error) {
	return DB.ListCollectionNames(context.TODO(), bson.D{})
}
