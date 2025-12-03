package banner

import (
	"errors"

	"github.com/madr/backend/internal/domain/banner"
	"github.com/madr/backend/pkg/database"
	"gorm.io/gorm"
)

// Repository defines the interface for banner repository
type Repository interface {
	Create(bnr *banner.Banner) error
	GetByID(id uint) (*banner.Banner, error)
	GetAll(limit, offset int) ([]banner.Banner, int64, error)
	Update(bnr *banner.Banner) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new banner repository
func NewRepository() Repository {
	return &repository{
		db: database.GetDB(),
	}
}

// Create creates a new banner
func (r *repository) Create(bnr *banner.Banner) error {
	if err := r.db.Create(bnr).Error; err != nil {
		return err
	}
	return nil
}

// GetByID retrieves a banner by ID
func (r *repository) GetByID(id uint) (*banner.Banner, error) {
	var bnr banner.Banner
	if err := r.db.First(&bnr, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("banner not found")
		}
		return nil, err
	}
	return &bnr, nil
}

// GetAll retrieves all banners with pagination
func (r *repository) GetAll(limit, offset int) ([]banner.Banner, int64, error) {
	var banners []banner.Banner
	var total int64

	// Count total records
	if err := r.db.Model(&banner.Banner{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated records, ordered by created_at descending (newest first)
	if err := r.db.Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&banners).Error; err != nil {
		return nil, 0, err
	}

	return banners, total, nil
}

// Update updates an existing banner
func (r *repository) Update(bnr *banner.Banner) error {
	if err := r.db.Save(bnr).Error; err != nil {
		return err
	}
	return nil
}

// Delete soft deletes a banner
func (r *repository) Delete(id uint) error {
	if err := r.db.Delete(&banner.Banner{}, id).Error; err != nil {
		return err
	}
	return nil
}

