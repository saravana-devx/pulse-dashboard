package auth

import (
	"time"
)

type User struct {
	ID           string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v7()"`
	Email        string    `gorm:"uniqueIndex;not null"`
	PasswordHash string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"default:now()"`
}

func (User) TableName() string {
	return "users"
}

type RefreshToken struct {
	ID        string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v7()"`
	UserID    string    `gorm:"not null"`
	TokenHash string    `gorm:"uniqueIndex;not null"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:now()"`
}

func (RefreshToken) TableName() string {
	return "refresh_tokens"
}
