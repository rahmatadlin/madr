package kajian

import (
	"errors"

	"github.com/madr/backend/internal/domain/kajian"
	"github.com/madr/backend/pkg/database"
	"gorm.io/gorm"
)

// Repository defines the interface for kajian repository
type Repository interface {
	CreateOrUpdate(k *kajian.Kajian) error
	GetByVideoID(videoID string) (*kajian.Kajian, error)
	GetAll(limit, offset int) ([]kajian.Kajian, int64, error)
	GetByID(id uint) (*kajian.Kajian, error)
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new kajian repository
func NewRepository() Repository {
	return &repository{db: database.GetDB()}
}

// CreateOrUpdate creates a new kajian or updates if video_id exists
func (r *repository) CreateOrUpdate(k *kajian.Kajian) error {
	var existing kajian.Kajian
	err := r.db.Where("video_id = ?", k.VideoID).First(&existing).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return r.db.Create(k).Error
	}
	if err != nil {
		return err
	}
	k.ID = existing.ID
	return r.db.Save(k).Error
}

// GetByVideoID retrieves a kajian by video_id
func (r *repository) GetByVideoID(videoID string) (*kajian.Kajian, error) {
	var k kajian.Kajian
	if err := r.db.Where("video_id = ?", videoID).First(&k).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("kajian not found")
		}
		return nil, err
	}
	return &k, nil
}

// GetAll retrieves all kajian with pagination, ordered by published_at descending
func (r *repository) GetAll(limit, offset int) ([]kajian.Kajian, int64, error) {
	var list []kajian.Kajian
	var total int64
	if err := r.db.Model(&kajian.Kajian{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Order("published_at DESC").Limit(limit).Offset(offset).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

// GetByID retrieves a kajian by ID
func (r *repository) GetByID(id uint) (*kajian.Kajian, error) {
	var k kajian.Kajian
	if err := r.db.First(&k, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("kajian not found")
		}
		return nil, err
	}
	return &k, nil
}

// Delete soft deletes a kajian
func (r *repository) Delete(id uint) error {
	return r.db.Delete(&kajian.Kajian{}, id).Error
}
