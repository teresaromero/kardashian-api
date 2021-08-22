package main

import (
	"kardashian_api/config"
	"kardashian_api/controllers"
	"kardashian_api/database"
	"kardashian_api/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	database.Connect()

}

func main() {
	router := gin.Default()
	api := router.Group("/api")
	v1 := api.Group("/v1")
	collections := v1.Group("collections")
	collections.Use(middlewares.ValidateCollection())
	{
		collections.GET("/:collection", controllers.Collection)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})

	router.Run(":5000")
}
