package about

import (
	"errors"

	aboutDomain "github.com/madr/backend/internal/domain/about"
	"github.com/madr/backend/pkg/database"
	"gorm.io/gorm"
)

// Repository defines the interface for about repository.
type Repository interface {
	GetLatest() (*aboutDomain.About, error)
	Create(abt *aboutDomain.About) error
	Update(abt *aboutDomain.About) error
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new about repository.
func NewRepository() Repository {
	return &repository{
		db: database.GetDB(),
	}
}

// GetLatest returns the most recently updated about record.
func (r *repository) GetLatest() (*aboutDomain.About, error) {
	var abt aboutDomain.About
	if err := r.db.Order("updated_at DESC").First(&abt).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &abt, nil
}

// Create inserts a new about record.
func (r *repository) Create(abt *aboutDomain.About) error {
	return r.db.Create(abt).Error
}

// Update updates an existing about record.
func (r *repository) Update(abt *aboutDomain.About) error {
	return r.db.Save(abt).Error
}
