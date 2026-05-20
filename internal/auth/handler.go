package auth

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var body CreateUserRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	result, err := h.svc.CreateUserSerive(ctx, &body)
	if err != nil {
		switch {
		case errors.Is(err, ErrEmailExists):
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		case errors.Is(err, ErrInvalidEmail):
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		case errors.As(err, new(*WeakPasswordError)):
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case errors.Is(err, ErrHashingPassword), errors.Is(err, ErrToCreateUser):
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":      "user created",
		"user":         result.User,
		"accessToken":  result.AccessToken,
		"refreshToken": result.RefreshToken,
	})
}

func (h *Handler) LoginUser(c *gin.Context) {
	var body LoginRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	loginResult, err := h.svc.LoginUserSerive(ctx, &body)
	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidEmail), errors.Is(err, ErrWrongPassword):
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Login Successful",
		"user":         loginResult.User,
		"accessToken":  loginResult.AccessToken,
		"refreshToken": loginResult.RefreshToken,
	})
}

func (h *Handler) RefreshAccessToken(c *gin.Context) {
	refreshToken, err := extractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	tokens, err := h.svc.RefreshAccessToken(ctx, refreshToken)
	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidRefreshToken),
			errors.Is(err, ErrRefreshTokenReused),
			errors.Is(err, ErrUserNotFound):
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "created new access and refresh token",
		"tokens": gin.H{
			"accessToken":  tokens.AccessToken,
			"refreshToken": tokens.RefreshToken,
		},
	})
}

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("missing authorization header")
	}
	parts := strings.SplitN(strings.TrimSpace(header), " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return "", errors.New("authorization header must be in form: Bearer <token>")
	}
	token := strings.TrimSpace(parts[1])
	if token == "" {
		return "", errors.New("bearer token is empty")
	}
	return token, nil
}
