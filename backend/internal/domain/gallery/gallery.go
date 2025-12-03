package gallery

import (
	"github.com/madr/backend/internal/domain/models"
)

// Gallery represents a gallery entity
type Gallery struct {
	models.BaseModel
	Title    string `gorm:"type:varchar(255);not null" json:"title" binding:"required"`
	ImageURL string `gorm:"type:varchar(500);not null" json:"image_url" binding:"required"`
}

// TableName specifies the table name for GORM
func (Gallery) TableName() string {
	return "gallery"
}

