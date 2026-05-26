package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"pulseDashboard/internal/auth"
	"pulseDashboard/internal/jobs"
	"pulseDashboard/internal/rabbitmq"
	"pulseDashboard/internal/redis"
)

func Register(r *gin.Engine, authHandler *auth.Handler, jobsHandler *jobs.Handler, jtiStore *auth.JTIStore, db *gorm.DB, rdb *redis.Redis, mq *rabbitmq.RabbitMQ) {
	RegisterHealthRoute(r, db, rdb, mq)
	RegisterAuthRoutes(r, authHandler, jtiStore)
	RegisterJobsRoutes(r, jobsHandler, jtiStore)
}
