package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"pulseDashboard/internal/auth"
	"pulseDashboard/internal/jobs"
	"pulseDashboard/internal/redis"
)

func Register(r *gin.Engine, authHandler *auth.Handler, jobsHandler *jobs.Handler, jtiStore *auth.JTIStore, db *gorm.DB, rdb *redis.Redis) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	RegisterHealthRoute(r, db, rdb)
	RegisterAuthRoutes(r, authHandler, jtiStore)
	RegisterJobsRoutes(r, jobsHandler, jtiStore)
}
