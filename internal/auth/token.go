package auth

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var secretKey = []byte("secret-key")

func CreateToken(id string, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"id":    id,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func HashToken(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func GenerateRefreshToken() (rawToken string, hashedToken string, err error) {
	bytes := make([]byte, 32)
	if _, err = rand.Read(bytes); err != nil {
		return "", "", err
	}
	rawToken = hex.EncodeToString(bytes)
	hashedToken, err = HashToken(rawToken)
	return rawToken, hashedToken, err
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
