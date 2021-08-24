package middlewares

import (
	"errors"
	"kardashian_api/database"
	"kardashian_api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateCollection() gin.HandlerFunc {
	return func(c *gin.Context) {
		collection := c.Param("collection")
		if collection != "" {
			isValid := database.ValidCollection(collection)

			if !isValid {
				utils.HandleError(c, http.StatusBadRequest, errors.New("invalid collection"))
			}
		}

		c.Next()
	}
}
