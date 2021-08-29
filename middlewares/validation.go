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
			list, _ := database.ListOfCollections()
			var isValid = false
			for _, col := range list {
				if col == collection {
					isValid = true
				}
			}

			if !isValid {
				err := http_errors.InvalidCollection(collection)
				utils.HandleHttpError(c, err)
			}
		}

		c.Next()
	}
}
