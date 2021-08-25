package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type routes struct {
	router *gin.Engine
}

func Routes() routes {
	r := routes{
		router: gin.Default(),
	}

	v1 := r.router.Group("/api/v1")

	r.collectionRoutes(v1)
	r.episodesRoutes(v1)

	r.router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
	})

	return r
}

func (r routes) Run(addr ...string) error {
	return r.router.Run()
}
