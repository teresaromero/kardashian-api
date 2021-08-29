package handlers

import (
	"kardashian_api/controllers"
	"kardashian_api/models"
	"kardashian_api/utils/request"
	"kardashian_api/utils/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllEpisodes(c *gin.Context) {
	p := request.GetContextValue(c.Request, "pagination").(*models.PaginationOpts)

	episodes, err := controllers.GetAllEpisodes(p)
	if err != nil {
		response.HttpError(c, err)
	} else {
		response.PageResponse(c, episodes)
	}
}

func GetEpisodeByNumber(c *gin.Context) {
	n, _ := strconv.Atoi(c.Param("number"))
	episode, err := controllers.GetEpisodeByNumber(n)
	if err != nil {
		response.HttpError(c, err)
	} else {
		response.SingleResponse(c, episode)
	}
}
