package auth

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           string         `gorm:"type:uuid;primaryKey;default:uuid_generate_v7()"`
	Email        string         `gorm:"uniqueIndex;not null"`
	PasswordHash string         `gorm:"not null"`
	CreatedAt    time.Time      `gorm:"default:now()"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (User) TableName() string {
	return "users"
}

type RefreshToken struct {
	ID        string     `gorm:"type:uuid;primaryKey;default:uuid_generate_v7()"`
	UserID    string     `gorm:"not null"`
	FamilyID  string     `gorm:"type:uuid;not null;default:uuid_generate_v7();index"`
	ParentID  *string    `gorm:"type:uuid"`
	TokenHash string     `gorm:"uniqueIndex;not null"`
	ExpiresAt time.Time  `gorm:"not null"`
	RevokedAt *time.Time
	CreatedAt time.Time `gorm:"default:now()"`
}

func (RefreshToken) TableName() string {
	return "refresh_tokens"
}
