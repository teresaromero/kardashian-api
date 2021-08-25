package routes

import (
	"kardashian_api/handlers"

	"github.com/gin-gonic/gin"
)

func (r routes) episodesRoutes(rg *gin.RouterGroup) {
	episodes := rg.Group("episodes")
	{
		episodes.GET("/", handlers.GetAllEpisodes)
		episodes.GET("/:number", handlers.GetEpisodeByNumber)

	}
}
