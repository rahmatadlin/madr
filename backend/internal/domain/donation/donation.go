package donation

import (
	"github.com/madr/backend/internal/domain/models"
)

// PaymentStatus represents payment status type
type PaymentStatus string

const (
	PaymentStatusPending PaymentStatus = "pending"
	PaymentStatusSuccess PaymentStatus = "success"
	PaymentStatusFailed  PaymentStatus = "failed"
)

// DonationCategoryInfo represents category information in donation response
type DonationCategoryInfo struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Donation represents a donation entity
type Donation struct {
	models.BaseModel
	CategoryID    uint          `gorm:"not null;index" json:"category_id" binding:"required"`
	DonorName     *string       `gorm:"type:varchar(255)" json:"donor_name"` // Nullable for anonymous
	Amount        float64       `gorm:"type:decimal(15,2);not null" json:"amount" binding:"required,gt=0"`
	Message       string        `gorm:"type:text" json:"message"`
	PaymentStatus PaymentStatus `gorm:"type:varchar(20);default:'pending'" json:"payment_status"`
	Category      *DonationCategoryInfo `gorm:"-" json:"category,omitempty"` // Will be loaded via Preload
}

