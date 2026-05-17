package auth

import (
	"errors"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CheckEmailExists(email string) (bool, error) {
	var user UserModel
	err := r.DB.Where("email = ?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *UserRepository) CreateUser(user *UserModel) error {
	return r.DB.Create(user).Error
}
