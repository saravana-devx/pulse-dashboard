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
	Router     *gin.Engine
	DB         *gorm.DB
	Redis      *redis.Redis
	TokenStore *auth.TokenStore
}

func New() (*App, error) {
	db, err := database.ConnectDB()
	if err != nil {
		return nil, err
	}
	rdb := redis.NewRedis()

	userRepo := auth.NewUserRepository(db)
	tokenStore := auth.NewTokenStore(rdb)
	authService := auth.NewService(userRepo)
	authHandler := auth.NewHandler(authService)

	router := gin.Default()
	routes.Register(router, authHandler)

	return &App{Router: router, DB: db, Redis: rdb, TokenStore: tokenStore}, nil
}
