package banner

import (
	"github.com/madr/backend/internal/domain/models"
)

// BannerType represents the type of banner media
type BannerType string

const (
	BannerTypeImage BannerType = "image"
	BannerTypeVideo BannerType = "video"
)

// Banner represents a banner entity
type Banner struct {
	models.BaseModel
	Title    string     `gorm:"type:varchar(255);not null" json:"title" binding:"required"`
	MediaURL string     `gorm:"type:varchar(500);not null" json:"media_url" binding:"required"`
	Type     BannerType `gorm:"type:varchar(20);not null" json:"type" binding:"required,oneof=image video"`
}

// TableName specifies the table name for GORM
func (Banner) TableName() string {
	return "banners"
}

