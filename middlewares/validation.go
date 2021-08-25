package middlewares

import (
	"kardashian_api/custom_errors"
	"kardashian_api/database"
	"kardashian_api/utils"

	"github.com/gin-gonic/gin"
)

func ValidateCollection() gin.HandlerFunc {
	return func(c *gin.Context) {
		collection := c.Param("collection")
		if collection != "" {
			isValid := database.ValidCollection(collection)

			if !isValid {
				err := custom_errors.InvalidCollection(collection)
				utils.HandleHttpError(c, err)
			}
		}

		c.Next()
	}
}
