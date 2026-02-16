package kajian

import (
	"time"

	"github.com/madr/backend/internal/domain/models"
)

// Kajian represents a synced YouTube video from the channel
type Kajian struct {
	models.BaseModel
	VideoID      string    `gorm:"type:varchar(20);not null;uniqueIndex" json:"video_id"`
	Title        string    `gorm:"type:varchar(255);not null" json:"title"`
	Description  string    `gorm:"type:text" json:"description"`
	PublishedAt  time.Time `gorm:"type:timestamp;not null" json:"published_at"`
	ThumbnailURL string    `gorm:"type:varchar(512)" json:"thumbnail_url"`
	YoutubeURL   string    `gorm:"type:varchar(512);not null" json:"youtube_url"`
	ChannelTitle string    `gorm:"type:varchar(255)" json:"channel_title"`
}

// TableName specifies the table name for GORM
func (Kajian) TableName() string {
	return "kajian"
}
