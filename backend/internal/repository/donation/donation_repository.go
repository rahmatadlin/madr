package donation

import (
	"errors"

	donationDomain "github.com/madr/backend/internal/domain/donation"
	"github.com/madr/backend/pkg/database"
	"gorm.io/gorm"
)

// Repository defines the interface for donation repository
type Repository interface {
	Create(don *donationDomain.Donation) error
	GetByID(id uint) (*donationDomain.Donation, error)
	GetAll(limit, offset int, status *donationDomain.PaymentStatus) ([]donationDomain.Donation, int64, error)
	Update(don *donationDomain.Donation) error
	Delete(id uint) error
	GetTotalAmount(status *donationDomain.PaymentStatus) (float64, error)
	GetTotalTransactions(status *donationDomain.PaymentStatus) (int64, error)
	GetAmountPerCategory(status *donationDomain.PaymentStatus) ([]CategoryAmount, error)
}

// CategoryAmount represents donation amount per category
type CategoryAmount struct {
	CategoryID   uint    `json:"category_id"`
	CategoryName string  `json:"category_name"`
	Amount       float64 `json:"amount"`
}

type repository struct {
	db *gorm.DB
}

// NewRepository creates a new donation repository
func NewRepository() Repository {
	return &repository{
		db: database.GetDB(),
	}
}

// Create creates a new donation
func (r *repository) Create(don *donationDomain.Donation) error {
	if err := r.db.Create(don).Error; err != nil {
		return err
	}
	return nil
}

// GetByID retrieves a donation by ID with category
func (r *repository) GetByID(id uint) (*donationDomain.Donation, error) {
	var don donationDomain.Donation
	if err := r.db.Preload("Category").First(&don, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("donation not found")
		}
		return nil, err
	}
	return &don, nil
}

// GetAll retrieves all donations with pagination and optional status filter
func (r *repository) GetAll(limit, offset int, status *donationDomain.PaymentStatus) ([]donationDomain.Donation, int64, error) {
	var donations []donationDomain.Donation
	var total int64

	query := r.db.Model(&donationDomain.Donation{})
	if status != nil {
		query = query.Where("payment_status = ?", *status)
	}

	// Count total records
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated records
	if err := query.Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&donations).Error; err != nil {
		return nil, 0, err
	}

	// Load categories for all donations
	categoryIDs := make([]uint, 0, len(donations))
	for _, d := range donations {
		categoryIDs = append(categoryIDs, d.CategoryID)
	}
	
	if len(categoryIDs) > 0 {
		var categories []struct {
			ID          uint
			Name        string
			Description string
		}
		r.db.Table("donation_categories").
			Select("id, name, description").
			Where("id IN ?", categoryIDs).
			Find(&categories)
		
		// Map categories to donations
		categoryMap := make(map[uint]*donationDomain.DonationCategoryInfo)
		for _, c := range categories {
			categoryMap[c.ID] = &donationDomain.DonationCategoryInfo{
				ID:          c.ID,
				Name:        c.Name,
				Description: c.Description,
			}
		}
		
		for i := range donations {
			if cat, ok := categoryMap[donations[i].CategoryID]; ok {
				donations[i].Category = cat
			}
		}
	}

	return donations, total, nil
}

// Update updates an existing donation
func (r *repository) Update(don *donationDomain.Donation) error {
	if err := r.db.Save(don).Error; err != nil {
		return err
	}
	return nil
}

// Delete soft deletes a donation
func (r *repository) Delete(id uint) error {
	if err := r.db.Delete(&donationDomain.Donation{}, id).Error; err != nil {
		return err
	}
	return nil
}

// GetTotalAmount calculates total donation amount (only success status by default)
func (r *repository) GetTotalAmount(status *donationDomain.PaymentStatus) (float64, error) {
	var total float64
	query := r.db.Model(&donationDomain.Donation{})
	
	if status != nil {
		query = query.Where("payment_status = ?", *status)
	} else {
		// Default to success only
		query = query.Where("payment_status = ?", donationDomain.PaymentStatusSuccess)
	}

	if err := query.Select("COALESCE(SUM(amount), 0)").Scan(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

// GetTotalTransactions counts total donation transactions
func (r *repository) GetTotalTransactions(status *donationDomain.PaymentStatus) (int64, error) {
	var total int64
	query := r.db.Model(&donationDomain.Donation{})
	
	if status != nil {
		query = query.Where("payment_status = ?", *status)
	} else {
		// Default to success only
		query = query.Where("payment_status = ?", donationDomain.PaymentStatusSuccess)
	}

	if err := query.Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

// GetAmountPerCategory calculates donation amount per category (optimized SQL)
func (r *repository) GetAmountPerCategory(status *donationDomain.PaymentStatus) ([]CategoryAmount, error) {
	var results []CategoryAmount
	
	query := r.db.Model(&donationDomain.Donation{}).
		Select(`
			donations.category_id,
			donation_categories.name as category_name,
			COALESCE(SUM(donations.amount), 0) as amount
		`).
		Joins("LEFT JOIN donation_categories ON donations.category_id = donation_categories.id").
		Group("donations.category_id, donation_categories.name")
	
	if status != nil {
		query = query.Where("donations.payment_status = ?", *status)
	} else {
		// Default to success only
		query = query.Where("donations.payment_status = ?", donationDomain.PaymentStatusSuccess)
	}

	if err := query.Order("amount DESC").Scan(&results).Error; err != nil {
		return nil, err
	}
	
	return results, nil
}

