package main

import (
	"kardashian_api/config"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
}

func main() {
	router := gin.Default()

	router.Run(":5000")
}
