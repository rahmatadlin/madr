package gallery

import (
	"errors"

	galleryDomain "github.com/madr/backend/internal/domain/gallery"
	galleryRepo "github.com/madr/backend/internal/repository/gallery"
	"github.com/madr/backend/pkg/logger"
)

// UseCase defines the interface for gallery use case
type UseCase interface {
	Create(req *CreateRequest) (*galleryDomain.Gallery, error)
	GetAll(limit, offset int) (*GetAllResponse, error)
	Delete(id uint) error
}

// CreateRequest represents the request to create a gallery item
type CreateRequest struct {
	Title    string `json:"title" binding:"required,min=3,max=255"`
	ImageURL string `json:"image_url" binding:"required"`
}

// GetAllResponse represents the response for getting all gallery items
type GetAllResponse struct {
	Data       []galleryDomain.Gallery `json:"data"`
	Total      int64                  `json:"total"`
	Limit      int                    `json:"limit"`
	Offset     int                    `json:"offset"`
	TotalPages int                    `json:"total_pages"`
}

type useCase struct {
	repo galleryRepo.Repository
}

// NewUseCase creates a new gallery use case
func NewUseCase(repo galleryRepo.Repository) UseCase {
	return &useCase{
		repo: repo,
	}
}

// Create creates a new gallery item
func (uc *useCase) Create(req *CreateRequest) (*galleryDomain.Gallery, error) {
	gal := &galleryDomain.Gallery{
		Title:    req.Title,
		ImageURL: req.ImageURL,
	}

	if err := uc.repo.Create(gal); err != nil {
		logger.Error().Err(err).Msg("Failed to create gallery item")
		return nil, errors.New("failed to create gallery item")
	}

	logger.Info().
		Uint("id", gal.ID).
		Str("title", gal.Title).
		Msg("Gallery item created successfully")

	return gal, nil
}

// GetAll retrieves all gallery items with pagination
func (uc *useCase) GetAll(limit, offset int) (*GetAllResponse, error) {
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

	items, total, err := uc.repo.GetAll(limit, offset)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get gallery items")
		return nil, errors.New("failed to get gallery items")
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return &GetAllResponse{
		Data:       items,
		Total:      total,
		Limit:      limit,
		Offset:     offset,
		TotalPages: totalPages,
	}, nil
}

// Delete deletes a gallery item
func (uc *useCase) Delete(id uint) error {
	if err := uc.repo.Delete(id); err != nil {
		logger.Error().Err(err).Uint("id", id).Msg("Failed to delete gallery item")
		return errors.New("failed to delete gallery item")
	}

	logger.Info().Uint("id", id).Msg("Gallery item deleted successfully")
	return nil
}

