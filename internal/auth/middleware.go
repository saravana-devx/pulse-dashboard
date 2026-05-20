package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	ctxUserID = "userID"
	ctxJTI    = "jti"
	ctxExp    = "exp"
)

func RequireAuth(jtiStore *JTIStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := extractBearerToken(c.GetHeader("Authorization"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		claims, err := ParseAccessToken(token)
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

		c.Set(ctxUserID, claims.UserID)
		c.Set(ctxJTI, claims.JTI)
		c.Set(ctxExp, claims.ExpiresAt)
		c.Next()
	}
}

func UserIDFromContext(c *gin.Context) (string, bool) {
	v, ok := c.Get(ctxUserID)
	if !ok {
		return "", false
	}
	s, ok := v.(string)
	return s, ok
}

func JTIFromContext(c *gin.Context) (string, bool) {
	v, ok := c.Get(ctxJTI)
	if !ok {
		return "", false
	}
	s, ok := v.(string)
	return s, ok
}

func ExpFromContext(c *gin.Context) (time.Time, bool) {
	v, ok := c.Get(ctxExp)
	if !ok {
		return time.Time{}, false
	}
	t, ok := v.(time.Time)
	return t, ok
}
