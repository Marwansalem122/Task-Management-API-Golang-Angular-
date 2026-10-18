package routes

import (
	"backend/internal/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(authHandler *handlers.AuthHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	auth := r.Group("/api/auth")
	{
		auth.POST("/register", authHandler.Register)
	}
	return r
}
