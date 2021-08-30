package config

import (
	"github.com/joho/godotenv"
	"os"
)

var MongoURI = "mongodb://localhost:27017"
var MongoDBNAME = ""
var PORT = "5000"

func LoadEnv(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}
	MongoURI = os.Getenv("MONGO_URI")
	MongoDBNAME = os.Getenv("MONGO_DBNAME")
	PORT = os.Getenv("PORT")
	return nil
}
