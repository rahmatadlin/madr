package user

import (
	"time"

	"github.com/madr/backend/internal/domain/models"
)

// UserRole represents user role type
type UserRole string

const (
	RoleAdmin UserRole = "admin"
	RoleUser  UserRole = "user"
)

// User represents a user entity
type User struct {
	models.BaseModel
	Username string   `gorm:"type:varchar(100);uniqueIndex;not null" json:"username" binding:"required"`
	Email    string   `gorm:"type:varchar(255);uniqueIndex;not null" json:"email" binding:"required,email"`
	Password string   `gorm:"type:varchar(255);not null" json:"-"` // Hidden from JSON
	Name     string   `gorm:"type:varchar(255)" json:"name"`
	Role     UserRole `gorm:"type:varchar(20);default:'user'" json:"role"`
	IsActive bool     `gorm:"default:true" json:"is_active"`
	LastLogin *time.Time `gorm:"type:timestamp" json:"last_login,omitempty"`
}

// TableName specifies the table name for GORM
func (User) TableName() string {
	return "users"
}

