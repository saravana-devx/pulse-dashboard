package auth

import (
	"fmt"
	"time"
)

type Service struct {
	repo *UserRepository
}

func NewService(repo *UserRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(req *CreateUserRequest) (*SignupResult, error) {
	if !validateEmail(req.Email) {
		return nil, ErrInvalidEmail
	}

	if err := validatePassword(req.Password); err != nil {
		return nil, wrapWeakPassword(err)
	}

	exists, err := s.repo.CheckEmailExists(req.Email)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInternal, err)
	}
	if exists {
		return nil, ErrEmailExists
	}

	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrHashingPassword, err)
	}

	rawRefresh, hashedRefresh, err := GenerateRefreshToken()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrRefreshTokenGenerate, err)
	}

	user := &User{
		Email:        req.Email,
		PasswordHash: hashedPassword,
	}
	refresh := &RefreshToken{
		TokenHash: hashedRefresh,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	createdUser, err := s.repo.CreateUserWithRefreshToken(user, refresh)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrToCreateUser, err)
	}

	accessToken, err := CreateToken(createdUser.ID, createdUser.Email)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrAccessTokenGenerate, err)
	}

	return &SignupResult{
		User:         createdUser,
		AccessToken:  accessToken,
		RefreshToken: rawRefresh,
	}, nil
}
