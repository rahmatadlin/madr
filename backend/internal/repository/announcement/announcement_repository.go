package announcement

import (
	"errors"

	"github.com/madr/backend/internal/domain/announcement"
	"github.com/madr/backend/pkg/database"
	"gorm.io/gorm"
)

// Repository defines the interface for announcement repository
type Repository interface {
	Create(ann *announcement.Announcement) error
	GetByID(id uint) (*announcement.Announcement, error)
	GetAll(limit, offset int) ([]announcement.Announcement, int64, error)
	GetPublished(limit, offset int) ([]announcement.Announcement, int64, error)
	Update(ann *announcement.Announcement) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new announcement repository
func NewRepository() Repository {
	return &repository{
		db: database.GetDB(),
	}
}

// Create creates a new announcement
func (r *repository) Create(ann *announcement.Announcement) error {
	if err := r.db.Create(ann).Error; err != nil {
		return err
	}
	return nil
}

// GetByID retrieves an announcement by ID
func (r *repository) GetByID(id uint) (*announcement.Announcement, error) {
	var ann announcement.Announcement
	if err := r.db.First(&ann, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("announcement not found")
		}
		return nil, err
	}
	return &ann, nil
}

// GetAll retrieves all announcements with pagination
func (r *repository) GetAll(limit, offset int) ([]announcement.Announcement, int64, error) {
	var announcements []announcement.Announcement
	var total int64

	// Count total records
	if err := r.db.Model(&announcement.Announcement{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated records
	if err := r.db.Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&announcements).Error; err != nil {
		return nil, 0, err
	}

	return announcements, total, nil
}

// GetPublished retrieves only published announcements with pagination
func (r *repository) GetPublished(limit, offset int) ([]announcement.Announcement, int64, error) {
	var announcements []announcement.Announcement
	var total int64

	// Count published records
	if err := r.db.Model(&announcement.Announcement{}).
		Where("is_published = ?", true).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated published records
	if err := r.db.Where("is_published = ?", true).
		Order("published_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&announcements).Error; err != nil {
		return nil, 0, err
	}

	return announcements, total, nil
}

// Update updates an existing announcement
func (r *repository) Update(ann *announcement.Announcement) error {
	if err := r.db.Save(ann).Error; err != nil {
		return err
	}
	return nil
}

// Delete soft deletes an announcement
func (r *repository) Delete(id uint) error {
	if err := r.db.Delete(&announcement.Announcement{}, id).Error; err != nil {
		return err
	}
	return nil
}

