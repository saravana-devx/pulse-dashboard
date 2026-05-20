package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

const (
	AccessTokenTTL  = 24 * time.Hour
	RefreshTokenTTL = 7 * 24 * time.Hour
)

var (
	jwtSecretOnce sync.Once
	jwtSecret     []byte
)

// getJWTSecret is lazy because viper is loaded inside main() — a package-level
// initializer would run before .env is read.
func getJWTSecret() []byte {
	jwtSecretOnce.Do(func() {
		s := viper.GetString("ACCESS_SECRET")
		if s == "" {
			panic("ACCESS_SECRET is not set in config")
		}
		jwtSecret = []byte(s)
	})
	return jwtSecret
}

func CreateToken(userID string, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"sub":   userID,
			"exp":   time.Now().Add(AccessTokenTTL).Unix(),
			"iat":   time.Now().Unix(),
		})
	return token.SignedString(getJWTSecret())
}

// HashToken returns a deterministic SHA-256 hex digest of the token.
// Deterministic hashing lets us look up refresh-token rows by hash. Bcrypt is
// the wrong tool here — it's salted/random and is for low-entropy passwords.
// Refresh tokens already carry 256 bits of entropy from crypto/rand.
func HashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}

func GenerateRefreshToken() (rawToken string, hashedToken string, err error) {
	bytes := make([]byte, 32)
	if _, err = rand.Read(bytes); err != nil {
		return "", "", err
	}
	rawToken = hex.EncodeToString(bytes)
	hashedToken = HashToken(rawToken)
	return rawToken, hashedToken, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return getJWTSecret(), nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}
