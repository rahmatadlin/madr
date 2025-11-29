package refreshtoken

import (
	"time"

	"github.com/madr/backend/internal/domain/models"
)

// RefreshToken represents a refresh token entity
type RefreshToken struct {
	models.BaseModel
	Token     string    `gorm:"type:varchar(500);uniqueIndex;not null" json:"token"`
	UserID    uint      `gorm:"not null;index" json:"user_id"`
	ExpiresAt time.Time `gorm:"type:timestamp;not null" json:"expires_at"`
	IsRevoked bool      `gorm:"default:false" json:"is_revoked"`
	RevokedAt *time.Time `gorm:"type:timestamp" json:"revoked_at,omitempty"`
	UserAgent string    `gorm:"type:varchar(255)" json:"user_agent,omitempty"`
	IPAddress string    `gorm:"type:varchar(45)" json:"ip_address,omitempty"`
}

// TableName specifies the table name for GORM
func (RefreshToken) TableName() string {
	return "refresh_tokens"
}

// IsExpired checks if the token is expired
func (rt *RefreshToken) IsExpired() bool {
	return time.Now().After(rt.ExpiresAt)
}

// IsValid checks if the token is valid (not expired and not revoked)
func (rt *RefreshToken) IsValid() bool {
	return !rt.IsExpired() && !rt.IsRevoked
}

