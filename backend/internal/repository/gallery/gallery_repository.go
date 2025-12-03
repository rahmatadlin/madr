package gallery

import (
	"github.com/madr/backend/internal/domain/gallery"
	"github.com/madr/backend/pkg/database"
	"gorm.io/gorm"
)

// Repository defines the interface for gallery repository
type Repository interface {
	Create(gal *gallery.Gallery) error
	GetAll(limit, offset int) ([]gallery.Gallery, int64, error)
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new gallery repository
func NewRepository() Repository {
	return &repository{
		db: database.GetDB(),
	}
}

// Create creates a new gallery item
func (r *repository) Create(gal *gallery.Gallery) error {
	if err := r.db.Create(gal).Error; err != nil {
		return err
	}
	return nil
}

// GetAll retrieves all gallery items with pagination
func (r *repository) GetAll(limit, offset int) ([]gallery.Gallery, int64, error) {
	var items []gallery.Gallery
	var total int64

	// Count total records
	if err := r.db.Model(&gallery.Gallery{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated records, ordered by created_at descending (newest first)
	if err := r.db.Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&items).Error; err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

// Delete soft deletes a gallery item
func (r *repository) Delete(id uint) error {
	if err := r.db.Delete(&gallery.Gallery{}, id).Error; err != nil {
		return err
	}
	return nil
}

