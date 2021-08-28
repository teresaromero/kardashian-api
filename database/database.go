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
		MongoUri := config.GetConfig().MongoUri

		client, errClient := mongo.NewClient(options.Client().ApplyURI(MongoUri))
		if errClient != nil {
			fmt.Println(errClient)
			panic(errClient)
		}
		dbClient = client

	}

	ctx, cancel := Context()
	defer cancel()

	errConn := dbClient.Connect(ctx)
	if errConn != nil {
		panic(errConn)
	}

	errPing := dbClient.Ping(ctx, nil)
	if errPing != nil {
		panic(errPing)
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
	return db.ListCollectionNames(context.TODO(), bson.D{})
}
