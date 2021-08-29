package config

import (
	"github.com/joho/godotenv"
	"os"
)

var MongoURI = "mongodb://localhost:27017"
var MongoDBNAME = "kuwtk"
var PORT = "5000"

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	MongoURI = os.Getenv("MONGO_URI")
	MongoDBNAME = os.Getenv("MONGO_DBNAME")
	PORT = os.Getenv("PORT")
	return nil
}
