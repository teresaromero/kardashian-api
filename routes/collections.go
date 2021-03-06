package routes

import (
	"kardashian_api/handlers"
	"kardashian_api/middlewares"

	"github.com/gin-gonic/gin"
)

func collectionRoutes(rg *gin.RouterGroup) {
	collections := rg.Group("collections")
	collections.Use(middlewares.ValidateCollection())
	collections.Use(middlewares.Pagination("collections"))
	{
		collections.GET("/", handlers.GetAvailableCollection)
		collections.GET("/:collection", handlers.GetCollection)
	}
}
