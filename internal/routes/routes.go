package routes

import (
	"github.com/gin-gonic/gin"

	"pulseDashboard/internal/auth"
)

func Register(r *gin.Engine, authHandler *auth.Handler) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	RegisterAuthRoutes(r, authHandler)
}
