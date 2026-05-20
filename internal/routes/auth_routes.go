package routes

import (
	"github.com/gin-gonic/gin"
	"pulseDashboard/internal/auth"
)

func RegisterAuthRoutes(r *gin.Engine, h *auth.Handler, jtiStore *auth.JTIStore) {
	g := r.Group("/auth")
	{
		g.POST("/sign-up", h.CreateUser)
		g.POST("/login", h.LoginUser)
		g.POST("/refresh", h.RefreshAccessToken)
		g.POST("/logout", auth.RequireAuth(jtiStore), h.Logout)
	}
}
