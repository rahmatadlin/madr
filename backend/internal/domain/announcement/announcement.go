package announcement

import (
	"time"

	"github.com/madr/backend/internal/domain/models"
)

// Announcement represents an announcement entity
type Announcement struct {
	models.BaseModel
	Title       string    `gorm:"type:varchar(255);not null" json:"title" binding:"required"`
	Content     string    `gorm:"type:text;not null" json:"content" binding:"required"`
	IsPublished bool      `gorm:"default:false" json:"is_published"`
	PublishedAt *time.Time `gorm:"type:timestamp" json:"published_at,omitempty"`
	Author      string    `gorm:"type:varchar(100)" json:"author"`
}

// TableName specifies the table name for GORM
func (Announcement) TableName() string {
	return "announcements"
}

// BeforeCreate hook to set PublishedAt if IsPublished is true
func (a *Announcement) BeforeCreate() error {
	if a.IsPublished && a.PublishedAt == nil {
		now := time.Now()
		a.PublishedAt = &now
	}
	return nil
}

// BeforeUpdate hook to update PublishedAt when IsPublished changes
func (a *Announcement) BeforeUpdate() error {
	if a.IsPublished && a.PublishedAt == nil {
		now := time.Now()
		a.PublishedAt = &now
	}
	return nil
}

