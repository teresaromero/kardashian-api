package handlers

import (
	"kardashian_api/controllers"
	"kardashian_api/utils/response"

	"github.com/gin-gonic/gin"
)

func GetCollection(c *gin.Context) {
	tableName := c.Param("collection")
	items, err := controllers.Collection(tableName)
	if err != nil {
		response.HttpError(c, err)
	} else {
		response.PageResponse(c, items)
	}
}

func GetAvailableCollection(c *gin.Context) {
	rsp, err := controllers.AvailableCollections(c.Request.URL.String())
	if err != nil {
		response.HttpError(c, err)
	} else {
		response.SingleResponse(c, rsp)
	}
}
