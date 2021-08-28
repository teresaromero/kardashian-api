package utils

import (
	"kardashian_api/utils/http_errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleHttpError(c *gin.Context, err *http_errors.HttpError) {
	c.AbortWithStatusJSON(err.Status(), gin.H{"status": err.Status(), "message": err.Error()})
}

func HandleSingleResponse(c *gin.Context, data interface{}, err *http_errors.HttpError) {
	if err != nil {
		HandleHttpError(c, err)
	} else {
		c.JSON(http.StatusOK, data)
	}
}

func HandlePageResponse(c *gin.Context, data interface{}, err *http_errors.HttpError) {
	if err != nil {
		HandleHttpError(c, err)
	} else {
		page := pageResponse(c.Request, data)
		c.JSON(http.StatusOK, page)
	}
}
