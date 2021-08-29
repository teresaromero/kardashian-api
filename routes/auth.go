package routes

import (
	"github.com/gin-gonic/gin"
	"kardashian_api/handlers"
)

func authRoutes(rg *gin.RouterGroup) {
	{
		rg.GET("/callback", handlers.CallbackHandler)
		rg.GET("/login", handlers.LoginHandler)
	}
}
