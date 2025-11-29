package user

import (
	"errors"
	"time"

	"github.com/madr/backend/internal/domain/user"
	"github.com/madr/backend/pkg/database"
	"gorm.io/gorm"
)

// Repository defines the interface for user repository
type Repository interface {
	Create(usr *user.User) error
	GetByID(id uint) (*user.User, error)
	GetByUsername(username string) (*user.User, error)
	GetByEmail(email string) (*user.User, error)
	Update(usr *user.User) error
	UpdateLastLogin(id uint) error
	Delete(id uint) error
	ExistsByUsername(username string) (bool, error)
	ExistsByEmail(email string) (bool, error)
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new user repository
func NewRepository() Repository {
	return &repository{
		db: database.GetDB(),
	}
}

// Create creates a new user
func (r *repository) Create(usr *user.User) error {
	if err := r.db.Create(usr).Error; err != nil {
		return err
	}
	return nil
}

// GetByID retrieves a user by ID
func (r *repository) GetByID(id uint) (*user.User, error) {
	var usr user.User
	if err := r.db.First(&usr, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &usr, nil
}

// GetByUsername retrieves a user by username
func (r *repository) GetByUsername(username string) (*user.User, error) {
	var usr user.User
	if err := r.db.Where("username = ?", username).First(&usr).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &usr, nil
}

// GetByEmail retrieves a user by email
func (r *repository) GetByEmail(email string) (*user.User, error) {
	var usr user.User
	if err := r.db.Where("email = ?", email).First(&usr).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &usr, nil
}

// Update updates an existing user
func (r *repository) Update(usr *user.User) error {
	if err := r.db.Save(usr).Error; err != nil {
		return err
	}
	return nil
}

// UpdateLastLogin updates the last login timestamp
func (r *repository) UpdateLastLogin(id uint) error {
	now := time.Now()
	if err := r.db.Model(&user.User{}).Where("id = ?", id).Update("last_login", now).Error; err != nil {
		return err
	}
	return nil
}

// Delete soft deletes a user
func (r *repository) Delete(id uint) error {
	if err := r.db.Delete(&user.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

// ExistsByUsername checks if a user exists by username
func (r *repository) ExistsByUsername(username string) (bool, error) {
	var count int64
	if err := r.db.Model(&user.User{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// ExistsByEmail checks if a user exists by email
func (r *repository) ExistsByEmail(email string) (bool, error) {
	var count int64
	if err := r.db.Model(&user.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

