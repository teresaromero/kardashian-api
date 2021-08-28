package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	MongoUri string
}

func GetConfig() Configuration {
	var configuration = Configuration{
		MongoUri: os.Getenv("MONGO_URI"),
	}
	return configuration
}

func LoadEnv() error {
	err := godotenv.Load()
	return err
}
