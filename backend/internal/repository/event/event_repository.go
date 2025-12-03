package event

import (
	"errors"

	"github.com/madr/backend/internal/domain/event"
	"github.com/madr/backend/pkg/database"
	"gorm.io/gorm"
)

// Repository defines the interface for event repository
type Repository interface {
	Create(evt *event.Event) error
	GetByID(id uint) (*event.Event, error)
	GetAll(limit, offset int) ([]event.Event, int64, error)
	Update(evt *event.Event) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new event repository
func NewRepository() Repository {
	return &repository{
		db: database.GetDB(),
	}
}

// Create creates a new event
func (r *repository) Create(evt *event.Event) error {
	if err := r.db.Create(evt).Error; err != nil {
		return err
	}
	return nil
}

// GetByID retrieves an event by ID
func (r *repository) GetByID(id uint) (*event.Event, error) {
	var evt event.Event
	if err := r.db.First(&evt, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("event not found")
		}
		return nil, err
	}
	return &evt, nil
}

// GetAll retrieves all events with pagination
func (r *repository) GetAll(limit, offset int) ([]event.Event, int64, error) {
	var events []event.Event
	var total int64

	// Count total records
	if err := r.db.Model(&event.Event{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated records, ordered by date ascending (upcoming events first)
	if err := r.db.Order("date ASC").
		Limit(limit).
		Offset(offset).
		Find(&events).Error; err != nil {
		return nil, 0, err
	}

	return events, total, nil
}

// Update updates an existing event
func (r *repository) Update(evt *event.Event) error {
	if err := r.db.Save(evt).Error; err != nil {
		return err
	}
	return nil
}

// Delete soft deletes an event
func (r *repository) Delete(id uint) error {
	if err := r.db.Delete(&event.Event{}, id).Error; err != nil {
		return err
	}
	return nil
}

