package refreshtoken

import (
	"errors"
	"time"

	"github.com/madr/backend/internal/domain/refreshtoken"
	"github.com/madr/backend/pkg/database"
	"gorm.io/gorm"
)

// Repository defines the interface for refresh token repository
type Repository interface {
	Create(rt *refreshtoken.RefreshToken) error
	GetByToken(token string) (*refreshtoken.RefreshToken, error)
	GetByUserID(userID uint) ([]refreshtoken.RefreshToken, error)
	Revoke(token string) error
	RevokeAllByUserID(userID uint) error
	DeleteExpired() error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new refresh token repository
func NewRepository() Repository {
	return &repository{
		db: database.GetDB(),
	}
}

// Create creates a new refresh token
func (r *repository) Create(rt *refreshtoken.RefreshToken) error {
	if err := r.db.Create(rt).Error; err != nil {
		return err
	}
	return nil
}

// GetByToken retrieves a refresh token by token string
func (r *repository) GetByToken(token string) (*refreshtoken.RefreshToken, error) {
	var rt refreshtoken.RefreshToken
	if err := r.db.Where("token = ?", token).First(&rt).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("refresh token not found")
		}
		return nil, err
	}
	return &rt, nil
}

// GetByUserID retrieves all refresh tokens for a user
func (r *repository) GetByUserID(userID uint) ([]refreshtoken.RefreshToken, error) {
	var tokens []refreshtoken.RefreshToken
	if err := r.db.Where("user_id = ?", userID).Find(&tokens).Error; err != nil {
		return nil, err
	}
	return tokens, nil
}

// Revoke revokes a refresh token
func (r *repository) Revoke(token string) error {
	now := time.Now()
	if err := r.db.Model(&refreshtoken.RefreshToken{}).
		Where("token = ?", token).
		Updates(map[string]interface{}{
			"is_revoked": true,
			"revoked_at": now,
		}).Error; err != nil {
		return err
	}
	return nil
}

// RevokeAllByUserID revokes all refresh tokens for a user
func (r *repository) RevokeAllByUserID(userID uint) error {
	now := time.Now()
	if err := r.db.Model(&refreshtoken.RefreshToken{}).
		Where("user_id = ? AND is_revoked = ?", userID, false).
		Updates(map[string]interface{}{
			"is_revoked": true,
			"revoked_at": now,
		}).Error; err != nil {
		return err
	}
	return nil
}

// DeleteExpired deletes expired refresh tokens
func (r *repository) DeleteExpired() error {
	now := time.Now()
	if err := r.db.Where("expires_at < ?", now).Delete(&refreshtoken.RefreshToken{}).Error; err != nil {
		return err
	}
	return nil
}

// Delete deletes a refresh token
func (r *repository) Delete(id uint) error {
	if err := r.db.Delete(&refreshtoken.RefreshToken{}, id).Error; err != nil {
		return err
	}
	return nil
}

