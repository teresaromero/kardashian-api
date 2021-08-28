package routes

import (
	"kardashian_api/handlers"
	"kardashian_api/middlewares"

	"github.com/gin-gonic/gin"
)

func episodesRoutes(rg *gin.RouterGroup) {
	episodes := rg.Group("episodes")
	episodes.Use(middlewares.Pagination("episodes"))
	{
		episodes.GET("/", handlers.GetAllEpisodes)
		episodes.GET("/:number", handlers.GetEpisodeByNumber)

	}
}
