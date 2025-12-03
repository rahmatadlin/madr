package donationcategory

import (
	"errors"

	donationCategoryDomain "github.com/madr/backend/internal/domain/donationcategory"
	donationCategoryRepo "github.com/madr/backend/internal/repository/donationcategory"
	"github.com/madr/backend/pkg/logger"
)

// UseCase defines the interface for donation category use case
type UseCase interface {
	Create(req *CreateRequest) (*donationCategoryDomain.DonationCategory, error)
	GetByID(id uint) (*donationCategoryDomain.DonationCategory, error)
	GetAll() ([]donationCategoryDomain.DonationCategory, error)
	Update(id uint, req *UpdateRequest) (*donationCategoryDomain.DonationCategory, error)
	Delete(id uint) error
}

// CreateRequest represents the request to create a donation category
type CreateRequest struct {
	Name        string `json:"name" binding:"required,min=3,max=255"`
	Description string `json:"description"`
}

// UpdateRequest represents the request to update a donation category
type UpdateRequest struct {
	Name        string `json:"name" binding:"min=3,max=255"`
	Description string `json:"description"`
}

type useCase struct {
	repo donationCategoryRepo.Repository
}

// NewUseCase creates a new donation category use case
func NewUseCase(repo donationCategoryRepo.Repository) UseCase {
	return &useCase{
		repo: repo,
	}
}

// Create creates a new donation category
func (uc *useCase) Create(req *CreateRequest) (*donationCategoryDomain.DonationCategory, error) {
	// Check if name already exists
	exists, err := uc.repo.ExistsByName(req.Name, 0)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to check category name existence")
		return nil, errors.New("failed to check category name")
	}
	if exists {
		return nil, errors.New("category name already exists")
	}

	cat := &donationCategoryDomain.DonationCategory{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := uc.repo.Create(cat); err != nil {
		logger.Error().Err(err).Msg("Failed to create donation category")
		return nil, errors.New("failed to create donation category")
	}

	logger.Info().
		Uint("id", cat.ID).
		Str("name", cat.Name).
		Msg("Donation category created successfully")

	return cat, nil
}

// GetByID retrieves a donation category by ID
func (uc *useCase) GetByID(id uint) (*donationCategoryDomain.DonationCategory, error) {
	cat, err := uc.repo.GetByID(id)
	if err != nil {
		logger.Error().Err(err).Uint("id", id).Msg("Failed to get donation category")
		return nil, err
	}
	return cat, nil
}

// GetAll retrieves all donation categories
func (uc *useCase) GetAll() ([]donationCategoryDomain.DonationCategory, error) {
	categories, err := uc.repo.GetAll()
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get donation categories")
		return nil, errors.New("failed to get donation categories")
	}
	return categories, nil
}

// Update updates an existing donation category
func (uc *useCase) Update(id uint, req *UpdateRequest) (*donationCategoryDomain.DonationCategory, error) {
	cat, err := uc.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Check if new name already exists (excluding current category)
	if req.Name != "" && req.Name != cat.Name {
		exists, err := uc.repo.ExistsByName(req.Name, id)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to check category name existence")
			return nil, errors.New("failed to check category name")
		}
		if exists {
			return nil, errors.New("category name already exists")
		}
		cat.Name = req.Name
	}

	if req.Description != "" {
		cat.Description = req.Description
	}

	if err := uc.repo.Update(cat); err != nil {
		logger.Error().Err(err).Uint("id", id).Msg("Failed to update donation category")
		return nil, errors.New("failed to update donation category")
	}

	logger.Info().
		Uint("id", cat.ID).
		Msg("Donation category updated successfully")

	return cat, nil
}

// Delete deletes a donation category
func (uc *useCase) Delete(id uint) error {
	if err := uc.repo.Delete(id); err != nil {
		logger.Error().Err(err).Uint("id", id).Msg("Failed to delete donation category")
		return errors.New("failed to delete donation category")
	}

	logger.Info().Uint("id", id).Msg("Donation category deleted successfully")
	return nil
}

