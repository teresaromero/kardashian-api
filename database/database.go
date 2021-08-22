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

		client, err_client := mongo.NewClient(options.Client().ApplyURI(MONGO_URI))
		if err_client != nil {
			fmt.Println(err_client)
			panic(err_client)
		}
		dbClient = client

	}

	ctx, cancel := Context()
	defer cancel()

	err_conn := dbClient.Connect(ctx)
	if err_conn != nil {
		panic(err_conn)
	}

	err_ping := dbClient.Ping(ctx, nil)
	if err_ping != nil {
		panic(err_ping)
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

func ValidCollection(coll string) bool {
	coll_lst, err_coll := db.ListCollectionNames(context.TODO(), bson.D{})

	if err_coll != nil {
		return false
	}
	for _, c := range coll_lst {
		if c == coll {
			return true
		}
	}
	return false
}
