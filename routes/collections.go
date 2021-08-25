package routes

import (
	"kardashian_api/controllers"
	"kardashian_api/middlewares"

	"github.com/gin-gonic/gin"
)

func (r routes) collectionRoutes(rg *gin.RouterGroup) {
	collections := rg.Group("collections")
	collections.Use(middlewares.ValidateCollection())
	{
		collections.GET("/:collection", controllers.Collection)
	}
}
