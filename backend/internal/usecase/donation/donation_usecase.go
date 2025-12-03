package donation

import (
	"errors"

	donationDomain "github.com/madr/backend/internal/domain/donation"
	donationRepo "github.com/madr/backend/internal/repository/donation"
	"github.com/madr/backend/pkg/logger"
)

// UseCase defines the interface for donation use case
type UseCase interface {
	Create(req *CreateRequest) (*donationDomain.Donation, error)
	GetByID(id uint) (*donationDomain.Donation, error)
	GetAll(limit, offset int, status *donationDomain.PaymentStatus) (*GetAllResponse, error)
	Update(id uint, req *UpdateRequest) (*donationDomain.Donation, error)
	Delete(id uint) error
	GetSummary() (*SummaryResponse, error)
}

// CreateRequest represents the request to create a donation
type CreateRequest struct {
	CategoryID    uint    `json:"category_id" binding:"required"`
	DonorName     *string `json:"donor_name"`
	Amount        float64 `json:"amount" binding:"required,gt=0"`
	Message       string  `json:"message"`
	PaymentStatus string  `json:"payment_status" binding:"omitempty,oneof=pending success failed"`
}

// UpdateRequest represents the request to update a donation
type UpdateRequest struct {
	CategoryID    *uint   `json:"category_id"`
	DonorName     *string `json:"donor_name"`
	Amount        *float64 `json:"amount" binding:"omitempty,gt=0"`
	Message       *string `json:"message"`
	PaymentStatus *string `json:"payment_status" binding:"omitempty,oneof=pending success failed"`
}

// GetAllResponse represents the response for getting all donations
type GetAllResponse struct {
	Data       []donationDomain.Donation `json:"data"`
	Total      int64                    `json:"total"`
	Limit      int                      `json:"limit"`
	Offset     int                      `json:"offset"`
	TotalPages int                      `json:"total_pages"`
}

// SummaryResponse represents the donation summary response
type SummaryResponse struct {
	TotalAmount      float64        `json:"total_amount"`
	TotalTransactions int64         `json:"total_transactions"`
	PerCategory      []CategorySummary `json:"per_category"`
}

// CategorySummary represents donation summary per category
type CategorySummary struct {
	CategoryID   uint    `json:"category_id"`
	CategoryName string  `json:"category"`
	Amount       float64 `json:"amount"`
}

type useCase struct {
	repo donationRepo.Repository
}

// NewUseCase creates a new donation use case
func NewUseCase(repo donationRepo.Repository) UseCase {
	return &useCase{
		repo: repo,
	}
}

// Create creates a new donation
func (uc *useCase) Create(req *CreateRequest) (*donationDomain.Donation, error) {
	// Set default payment status
	paymentStatus := donationDomain.PaymentStatusPending
	if req.PaymentStatus != "" {
		paymentStatus = donationDomain.PaymentStatus(req.PaymentStatus)
	}

	don := &donationDomain.Donation{
		CategoryID:    req.CategoryID,
		DonorName:     req.DonorName,
		Amount:        req.Amount,
		Message:       req.Message,
		PaymentStatus: paymentStatus,
	}

	if err := uc.repo.Create(don); err != nil {
		logger.Error().Err(err).Msg("Failed to create donation")
		return nil, errors.New("failed to create donation")
	}

	logger.Info().
		Uint("id", don.ID).
		Uint("category_id", don.CategoryID).
		Float64("amount", don.Amount).
		Str("payment_status", string(don.PaymentStatus)).
		Msg("Donation created successfully")

	return don, nil
}

// GetByID retrieves a donation by ID
func (uc *useCase) GetByID(id uint) (*donationDomain.Donation, error) {
	don, err := uc.repo.GetByID(id)
	if err != nil {
		logger.Error().Err(err).Uint("id", id).Msg("Failed to get donation")
		return nil, err
	}
	return don, nil
}

// GetAll retrieves all donations with pagination
func (uc *useCase) GetAll(limit, offset int, status *donationDomain.PaymentStatus) (*GetAllResponse, error) {
	// Validate pagination parameters
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	donations, total, err := uc.repo.GetAll(limit, offset, status)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get donations")
		return nil, errors.New("failed to get donations")
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return &GetAllResponse{
		Data:       donations,
		Total:      total,
		Limit:      limit,
		Offset:     offset,
		TotalPages: totalPages,
	}, nil
}

// Update updates an existing donation
func (uc *useCase) Update(id uint, req *UpdateRequest) (*donationDomain.Donation, error) {
	don, err := uc.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if req.CategoryID != nil {
		don.CategoryID = *req.CategoryID
	}
	if req.DonorName != nil {
		don.DonorName = req.DonorName
	}
	if req.Amount != nil {
		don.Amount = *req.Amount
	}
	if req.Message != nil {
		don.Message = *req.Message
	}
	if req.PaymentStatus != nil {
		don.PaymentStatus = donationDomain.PaymentStatus(*req.PaymentStatus)
	}

	if err := uc.repo.Update(don); err != nil {
		logger.Error().Err(err).Uint("id", id).Msg("Failed to update donation")
		return nil, errors.New("failed to update donation")
	}

	logger.Info().
		Uint("id", don.ID).
		Msg("Donation updated successfully")

	return don, nil
}

// Delete deletes a donation
func (uc *useCase) Delete(id uint) error {
	if err := uc.repo.Delete(id); err != nil {
		logger.Error().Err(err).Uint("id", id).Msg("Failed to delete donation")
		return errors.New("failed to delete donation")
	}

	logger.Info().Uint("id", id).Msg("Donation deleted successfully")
	return nil
}

// GetSummary calculates donation summary with optimized SQL
func (uc *useCase) GetSummary() (*SummaryResponse, error) {
	// Get total amount (only success payments)
	successStatus := donationDomain.PaymentStatusSuccess
	totalAmount, err := uc.repo.GetTotalAmount(&successStatus)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to calculate total amount")
		return nil, errors.New("failed to calculate total amount")
	}

	// Get total transactions (only success payments)
	totalTransactions, err := uc.repo.GetTotalTransactions(&successStatus)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to count total transactions")
		return nil, errors.New("failed to count total transactions")
	}

	// Get amount per category (only success payments)
	categoryAmounts, err := uc.repo.GetAmountPerCategory(&successStatus)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to calculate amount per category")
		return nil, errors.New("failed to calculate amount per category")
	}

	// Convert to summary format
	perCategory := make([]CategorySummary, len(categoryAmounts))
	for i, ca := range categoryAmounts {
		perCategory[i] = CategorySummary{
			CategoryID:   ca.CategoryID,
			CategoryName: ca.CategoryName,
			Amount:       ca.Amount,
		}
	}

	logger.Info().
		Float64("total_amount", totalAmount).
		Int64("total_transactions", totalTransactions).
		Msg("Donation summary calculated successfully")

	return &SummaryResponse{
		TotalAmount:      totalAmount,
		TotalTransactions: totalTransactions,
		PerCategory:      perCategory,
	}, nil
}

