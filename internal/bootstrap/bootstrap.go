package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"pulseDashboard/internal/auth"
	"pulseDashboard/internal/database"
	"pulseDashboard/internal/routes"
)

type App struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func New() (*App, error) {
	db, err := database.ConnectDB()
	if err != nil {
		return nil, err
	}

	// auth wiring
	userRepo := auth.NewUserRepository(db)
	authService := auth.NewService(userRepo)
	authHandler := auth.NewHandler(authService)

	router := gin.Default()
	routes.Register(router, authHandler)

	return &App{Router: router, DB: db}, nil
}
