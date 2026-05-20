package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"pulseDashboard/internal/auth"
	"pulseDashboard/internal/database"
	"pulseDashboard/internal/redis"
	"pulseDashboard/internal/routes"
)

type App struct {
	Router *gin.Engine
	DB     *gorm.DB
	Redis  *redis.Redis
}

func New() (*App, error) {
	db, err := database.ConnectDB()
	if err != nil {
		return nil, err
	}
	rdb := redis.NewRedis()

	userRepo := auth.NewUserRepository(db)
	jtiStore := auth.NewJTIStore(rdb)
	authService := auth.NewService(userRepo, jtiStore)
	authHandler := auth.NewHandler(authService)

	router := gin.Default()
	routes.Register(router, authHandler, jtiStore, db, rdb)

	return &App{Router: router, DB: db, Redis: rdb}, nil
}
