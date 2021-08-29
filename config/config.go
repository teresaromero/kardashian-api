package config

import (
	"github.com/joho/godotenv"
	"os"
)

var MongoURI = "mongodb://localhost:27017"
var MongoDBNAME = "kuwtk"
var PORT = "5000"
var Auth0ClientId = ""
var Auth0ClientSecret = ""
var Auth0RedirectUrl = ""
var Auth0Domain = ""

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	MongoURI = os.Getenv("MONGO_URI")
	MongoDBNAME = os.Getenv("MONGO_DBNAME")
	PORT = os.Getenv("PORT")
	Auth0ClientId = os.Getenv("AUTH0_CLIENT_ID")
	Auth0ClientSecret = os.Getenv("AUTH0_CLIENT_SECRET")
	Auth0RedirectUrl = os.Getenv("AUTH0_REDIRECT_URL")
	Auth0Domain = os.Getenv("AUTH0_DOMAIN")

	return nil
}
