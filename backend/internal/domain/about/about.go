package about

import (
	"github.com/madr/backend/internal/domain/models"
)

// About represents the about section content.
type About struct {
	models.BaseModel
	Title                 string `gorm:"type:varchar(255);not null" json:"title"`
	Subtitle              string `gorm:"type:varchar(255)" json:"subtitle"`
	Description           string `gorm:"type:text" json:"description"`
	AdditionalDescription string `gorm:"type:text" json:"additional_description"`
	ImageURL              string `gorm:"type:varchar(500)" json:"image_url"`
	YearsActive           int    `json:"years_active"`
	ActiveMembers         int    `json:"active_members"`
}

// TableName specifies the table name for GORM.
func (About) TableName() string {
	return "about"
}
