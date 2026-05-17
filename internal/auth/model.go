package auth

import (
	"time"
)

type UserModel struct {
	ID           string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v7()"`
	Email        string    `gorm:"uniqueIndex;not null"`
	PasswordHash string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"default:now()"`
}

func (UserModel) TableName() string {
	return "users"
}
