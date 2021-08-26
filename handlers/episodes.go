package handlers

import (
	"kardashian_api/controllers"
	"kardashian_api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllEpisodes(c *gin.Context) {
	p := utils.GetPaginationContext(c.Request)

	episodes, err := controllers.GetAllEpisodes(p)
	utils.HandlePageResponse(c, episodes, err)
}

func GetEpisodeByNumber(c *gin.Context) {
	n, _ := strconv.Atoi(c.Param("number"))
	episode, err := controllers.GetEpisodeByNumber(n)
	utils.HandleSingleResponse(c, episode, err)
}
