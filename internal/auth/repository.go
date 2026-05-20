package auth

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CheckEmailExists(ctx context.Context, email string) (bool, error) {
	var user User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *UserRepository) CreateUserWithRefreshToken(ctx context.Context, user *User, refreshToken *RefreshToken) (*User, error) {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		refreshToken.UserID = user.ID
		return tx.Create(refreshToken).Error
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUserPassword(ctx context.Context, email string) (*User, error) {
	var user User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, userID string) (*User, error) {
	var user User
	err := r.db.WithContext(ctx).Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateRefreshToken(ctx context.Context, rt *RefreshToken) error {
	return r.db.WithContext(ctx).Create(rt).Error
}

// FindRefreshTokenByHash looks up a row by hash regardless of status (active,
// expired, or revoked). Callers decide what to do based on the row's state —
// this is what enables reuse detection.
func (r *UserRepository) FindRefreshTokenByHash(ctx context.Context, tokenHash string) (*RefreshToken, error) {
	var rt RefreshToken
	err := r.db.WithContext(ctx).Where("token_hash = ?", tokenHash).First(&rt).Error
	if err != nil {
		return nil, err
	}
	return &rt, nil
}

// RotateRefreshToken atomically marks the old row as revoked and inserts a new
// row in the same family, linked via parent_id. Scoped to (id, user_id) so a
// caller can never rotate someone else's row.
func (r *UserRepository) RotateRefreshToken(ctx context.Context, old *RefreshToken, newHashedRefreshToken string) (*RefreshToken, error) {
	newToken := &RefreshToken{
		UserID:    old.UserID,
		FamilyID:  old.FamilyID,
		ParentID:  &old.ID,
		TokenHash: newHashedRefreshToken,
		ExpiresAt: time.Now().Add(RefreshTokenTTL),
	}

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		now := time.Now()
		res := tx.Model(&RefreshToken{}).
			Where("id = ? AND user_id = ? AND revoked_at IS NULL", old.ID, old.UserID).
			Update("revoked_at", now)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}
		return tx.Create(newToken).Error
	})
	if err != nil {
		return nil, err
	}
	return newToken, nil
}

// RevokeFamily revokes every active row in a token family. Used on reuse
// detection — if an old (already-revoked) token is replayed, we assume it's
// stolen and kill the entire chain so the attacker can't keep rotating.
// Other families for the same user (e.g. their other devices) are untouched.
func (r *UserRepository) RevokeFamily(ctx context.Context, familyID string) error {
	return r.db.WithContext(ctx).
		Model(&RefreshToken{}).
		Where("family_id = ? AND revoked_at IS NULL", familyID).
		Update("revoked_at", time.Now()).Error
}
