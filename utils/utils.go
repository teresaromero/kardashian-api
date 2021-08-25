package utils

import (
	errors "kardashian_api/custom_errors"

	"github.com/gin-gonic/gin"
)

func HandleHttpError(c *gin.Context, err *errors.HttpError) {
	c.AbortWithStatusJSON(err.Status(), gin.H{"status": err.Status(), "message": err.Error()})
}
