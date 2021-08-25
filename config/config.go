package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	MONGO_URI string
}

func GetConfig() Configuration {
	var configuration = Configuration{
		MONGO_URI: os.Getenv("MONGO_URI"),
	}
	return configuration
}

func LoadEnv() {
	err_env := godotenv.Load()
	if err_env != nil {
		log.Fatal("Error loading .env file")
	}
}
