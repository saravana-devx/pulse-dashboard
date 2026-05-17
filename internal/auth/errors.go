package auth

import (
	"errors"
)

var (
	ErrEmailExists     = errors.New("email already exists")
	ErrInvalidEmail    = errors.New("invalid email address")
	ErrWeakPassword    = errors.New("weak password")
	ErrInternal        = errors.New("internal error")
	ErrHashingPassword = errors.New("problem occured while hashing the password")
	ErrToCreateUser    = errors.New("failed to create user")
)

type WeakPasswordError struct{ Reason string }

func (e *WeakPasswordError) Error() string { return e.Reason }

func wrapWeaKPassword(err error) error {
	return &WeakPasswordError{Reason: err.Error()}
}
