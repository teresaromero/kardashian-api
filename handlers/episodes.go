package handlers

import (
	"kardashian_api/controllers"
	"kardashian_api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllEpisodes(c *gin.Context) {
	episodes, err := controllers.GetAllEpisodes()
	if err != nil {
		utils.HandleHttpError(c, err)
	} else {
		c.JSON(http.StatusOK, episodes)
	}
}

func GetEpisodeByNumber(c *gin.Context) {
	n, _ := strconv.Atoi(c.Param("number"))
	episode, err := controllers.GetEpisodeByNumber(n)
	if err != nil {
		utils.HandleHttpError(c, err)
	} else {
		c.JSON(http.StatusOK, episode)
	}
}
