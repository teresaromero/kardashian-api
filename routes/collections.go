package routes

import (
	"kardashian_api/handlers"
	"kardashian_api/middlewares"

	"github.com/gin-gonic/gin"
)

func (r routes) collectionRoutes(rg *gin.RouterGroup) {
	collections := rg.Group("collections")
	collections.Use(middlewares.ValidateCollection())
	{
		collections.GET("/", handlers.GetAvailableCollection)
		collections.GET("/:collection", handlers.GetCollection)
	}
}
