package auth

import (
	"errors"
)

var (
	ErrEmailExists          = errors.New("email already exists")
	ErrInvalidEmail         = errors.New("invalid email address")
	ErrWeakPassword         = errors.New("weak password")
	ErrInternal             = errors.New("internal error")
	ErrHashingPassword      = errors.New("problem occured while hashing the password")
	ErrToCreateUser         = errors.New("failed to create user")
	ErrAccessTokenGenerate  = errors.New("failed to generate access token")
	ErrRefreshTokenGenerate = errors.New("failed to generate refresh token")
	ErrRefreshTokenStore    = errors.New("failed to create refresh token entry in database")
)

type WeakPasswordError struct{ Reason string }

func (e *WeakPasswordError) Error() string { return e.Reason }

func wrapWeakPassword(err error) error {
	return &WeakPasswordError{Reason: err.Error()}
}
