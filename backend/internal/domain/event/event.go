package event

import (
	"time"

	"github.com/madr/backend/internal/domain/models"
)

// Event represents an event entity
type Event struct {
	models.BaseModel
	Title       string    `gorm:"type:varchar(255);not null" json:"title" binding:"required"`
	Description string    `gorm:"type:text" json:"description"`
	Date        time.Time `gorm:"type:timestamp;not null" json:"date" binding:"required"`
	Location    string    `gorm:"type:varchar(255)" json:"location"`
}

// TableName specifies the table name for GORM
func (Event) TableName() string {
	return "events"
}

