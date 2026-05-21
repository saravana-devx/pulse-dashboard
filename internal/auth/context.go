package auth

import (
	"time"

	"github.com/gin-gonic/gin"
)

const (
	ctxUserID = "userID"
	ctxJTI    = "jti"
	ctxExp    = "exp"
)

func SetUserID(c *gin.Context, userID string) { c.Set(ctxUserID, userID) }
func SetJTI(c *gin.Context, jti string)       { c.Set(ctxJTI, jti) }
func SetExp(c *gin.Context, exp time.Time)    { c.Set(ctxExp, exp) }

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
