package handlers

import (
	"kardashian_api/controllers"
	"kardashian_api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCollection(c *gin.Context) {
	tableName := c.Param("collection")
	items, err := controllers.Collection(tableName)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, items)
	}
}

func GetAvailableCollection(c *gin.Context) {
	rsp, err := controllers.AvailableCollections(c.Request.URL.String())
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, rsp)
	}
}
