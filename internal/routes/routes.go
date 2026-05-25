package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"pulseDashboard/internal/auth"
	"pulseDashboard/internal/httpx"
	"pulseDashboard/internal/jobs"
	"pulseDashboard/internal/redis"
)

func Register(r *gin.Engine, authHandler *auth.Handler, jobsHandler *jobs.Handler, jtiStore *auth.JTIStore, db *gorm.DB, rdb *redis.Redis) {
	r.GET("/ping", func(c *gin.Context) {
		httpx.Success(c, http.StatusOK, "pong", nil)
	})

	RegisterHealthRoute(r, db, rdb)
	RegisterAuthRoutes(r, authHandler, jtiStore)
	RegisterJobsRoutes(r, jobsHandler, jtiStore)
}
