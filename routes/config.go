package routes

import (
	"kardashian_api/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Router = gin.Default()

func LoadRoutes() {

	v1 := Router.Group("/api/v1")

	collectionRoutes(v1)
	episodesRoutes(v1)

	Router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
	})

}

func Run() error {
	err := Router.Run(":" + config.PORT)
	if err != nil {
		return err
	}
	return nil
}
