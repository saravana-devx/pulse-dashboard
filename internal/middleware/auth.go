package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"pulseDashboard/internal/auth"
)

func RequireAuth(jtiStore *auth.JTIStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := auth.ExtractBearerToken(c.GetHeader("Authorization"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		claims, err := auth.ParseAccessToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
		defer cancel()

		revoked, err := jtiStore.IsRevoked(ctx, claims.JTI)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "auth check failed"})
			return
		}
		if revoked {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token revoked"})
			return
		}

		auth.SetUserID(c, claims.UserID)
		auth.SetJTI(c, claims.JTI)
		auth.SetExp(c, claims.ExpiresAt)
		c.Next()
	}
}
