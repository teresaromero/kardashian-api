package main

import (
	"kardashian_api/config"
	"kardashian_api/database"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	database.Connect()

}

func main() {
	router := gin.Default()

	router.Run(":5000")
}
