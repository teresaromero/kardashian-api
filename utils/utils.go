package utils

import "github.com/gin-gonic/gin"

func HandleError(c *gin.Context, status int, err error) {
	c.AbortWithStatusJSON(status, gin.H{"status": status, "message": err.Error()})
}
