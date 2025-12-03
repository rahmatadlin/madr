package donationcategory

import (
	"github.com/madr/backend/internal/domain/models"
)

// DonationCategory represents a donation category entity
type DonationCategory struct {
	models.BaseModel
	Name        string `gorm:"type:varchar(255);not null;uniqueIndex" json:"name" binding:"required"`
	Description string `gorm:"type:text" json:"description"`
}

// TableName specifies the table name for GORM
func (DonationCategory) TableName() string {
	return "donation_categories"
}

