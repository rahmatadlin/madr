package donationcategory

import (
	"errors"

	"github.com/madr/backend/internal/domain/donationcategory"
	"github.com/madr/backend/pkg/database"
	"gorm.io/gorm"
)

// Repository defines the interface for donation category repository
type Repository interface {
	Create(cat *donationcategory.DonationCategory) error
	GetByID(id uint) (*donationcategory.DonationCategory, error)
	GetAll() ([]donationcategory.DonationCategory, error)
	Update(cat *donationcategory.DonationCategory) error
	Delete(id uint) error
	ExistsByName(name string, excludeID uint) (bool, error)
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new donation category repository
func NewRepository() Repository {
	return &repository{
		db: database.GetDB(),
	}
}

// Create creates a new donation category
func (r *repository) Create(cat *donationcategory.DonationCategory) error {
	if err := r.db.Create(cat).Error; err != nil {
		return err
	}
	return nil
}

// GetByID retrieves a donation category by ID
func (r *repository) GetByID(id uint) (*donationcategory.DonationCategory, error) {
	var cat donationcategory.DonationCategory
	if err := r.db.First(&cat, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("donation category not found")
		}
		return nil, err
	}
	return &cat, nil
}

// GetAll retrieves all donation categories
func (r *repository) GetAll() ([]donationcategory.DonationCategory, error) {
	var categories []donationcategory.DonationCategory
	if err := r.db.Order("name ASC").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

// Update updates an existing donation category
func (r *repository) Update(cat *donationcategory.DonationCategory) error {
	if err := r.db.Save(cat).Error; err != nil {
		return err
	}
	return nil
}

// Delete soft deletes a donation category
func (r *repository) Delete(id uint) error {
	if err := r.db.Delete(&donationcategory.DonationCategory{}, id).Error; err != nil {
		return err
	}
	return nil
}

// ExistsByName checks if a category with the given name exists
func (r *repository) ExistsByName(name string, excludeID uint) (bool, error) {
	var count int64
	query := r.db.Model(&donationcategory.DonationCategory{}).Where("name = ?", name)
	if excludeID > 0 {
		query = query.Where("id != ?", excludeID)
	}
	if err := query.Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

