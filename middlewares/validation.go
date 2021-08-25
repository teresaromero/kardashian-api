package middlewares

import (
	"kardashian_api/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateCollection() gin.HandlerFunc {
	return func(c *gin.Context) {
		collection := c.Param("collection")
		if collection != "" {
			isValid := database.ValidCollection(collection)

			if !isValid {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Collection Not Valid"})
			}
		}

		c.Next()
	}
}
