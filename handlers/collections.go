package handlers

import (
	"kardashian_api/controllers"
	"kardashian_api/utils"

	"github.com/gin-gonic/gin"
)

func GetCollection(c *gin.Context) {
	tableName := c.Param("collection")
	items, err := controllers.Collection(tableName)
	utils.HandlePageResponse(c, items, err)
}

func GetAvailableCollection(c *gin.Context) {
	rsp, err := controllers.AvailableCollections(c.Request.URL.String())
	utils.HandleResponse(c, rsp, err)
}
