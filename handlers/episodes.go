package handlers

import (
	"kardashian_api/repository"
	"kardashian_api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllEpisodes(c *gin.Context) {
	episodes, err := repository.GetAllEpisodes()
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, episodes)
	}
}

func GetEpisodeByNumber(c *gin.Context) {
	n, _ := strconv.Atoi(c.Param("number"))
	episode, err := repository.GetEpisodeByNumber(n)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, episode)
	}
}
