package middlewares

import (
	"kardashian_api/database"
	"kardashian_api/utils"
	"kardashian_api/utils/http_errors"

	"github.com/gin-gonic/gin"
)

func ValidateCollection() gin.HandlerFunc {
	return func(c *gin.Context) {
		collection := c.Param("collection")
		if collection != "" {
			isValid := database.ValidCollection(collection)

			if !isValid {
				err := http_errors.InvalidCollection(collection)
				utils.HandleHttpError(c, err)
			}
		}

		c.Next()
	}
}
