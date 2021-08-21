package database

import (
	"context"
	"fmt"
	"kardashian_api/config"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName string = "kuwtk"
)

var dbClient *mongo.Client
var db *mongo.Database

func Connect() {

	if dbClient == nil {
		MONGO_URI := config.GetConfig().MONGO_URI

		client, err := mongo.NewClient(options.Client().ApplyURI(MONGO_URI))
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		dbClient = client

	}

	ctx, cancel := Context()
	defer cancel()

	err := dbClient.Connect(ctx)
	if err != nil {
		panic(err)
	}

	err = dbClient.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	db = dbClient.Database(dbName)
}

func Use(tableName string) *mongo.Collection {
	coll := db.Collection(tableName)
	return coll
}

func Context() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}
