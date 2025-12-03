package banner

import (
	"errors"

	bannerDomain "github.com/madr/backend/internal/domain/banner"
	bannerRepo "github.com/madr/backend/internal/repository/banner"
	"github.com/madr/backend/pkg/logger"
)

// UseCase defines the interface for banner use case
type UseCase interface {
	Create(req *CreateRequest) (*bannerDomain.Banner, error)
	GetByID(id uint) (*bannerDomain.Banner, error)
	GetAll(limit, offset int) (*GetAllResponse, error)
	Update(id uint, req *UpdateRequest) (*bannerDomain.Banner, error)
	Delete(id uint) error
}

// CreateRequest represents the request to create a banner
type CreateRequest struct {
	Title    string `json:"title" binding:"required,min=3,max=255"`
	MediaURL string `json:"media_url" binding:"required"`
	Type     string `json:"type" binding:"required,oneof=image video"`
}

// UpdateRequest represents the request to update a banner
type UpdateRequest struct {
	Title    string `json:"title" binding:"min=3,max=255"`
	MediaURL string `json:"media_url"`
	Type     string `json:"type" binding:"oneof=image video"`
}

// GetAllResponse represents the response for getting all banners
type GetAllResponse struct {
	Data       []bannerDomain.Banner `json:"data"`
	Total      int64                `json:"total"`
	Limit      int                  `json:"limit"`
	Offset     int                  `json:"offset"`
	TotalPages int                  `json:"total_pages"`
}

type useCase struct {
	repo bannerRepo.Repository
}

// NewUseCase creates a new banner use case
func NewUseCase(repo bannerRepo.Repository) UseCase {
	return &useCase{
		repo: repo,
	}
}

// Create creates a new banner
func (uc *useCase) Create(req *CreateRequest) (*bannerDomain.Banner, error) {
	bnr := &bannerDomain.Banner{
		Title:    req.Title,
		MediaURL: req.MediaURL,
		Type:     bannerDomain.BannerType(req.Type),
	}

	if err := uc.repo.Create(bnr); err != nil {
		logger.Error().Err(err).Msg("Failed to create banner")
		return nil, errors.New("failed to create banner")
	}

	logger.Info().
		Uint("id", bnr.ID).
		Str("title", bnr.Title).
		Str("type", string(bnr.Type)).
		Msg("Banner created successfully")

	return bnr, nil
}

// GetByID retrieves a banner by ID
func (uc *useCase) GetByID(id uint) (*bannerDomain.Banner, error) {
	bnr, err := uc.repo.GetByID(id)
	if err != nil {
		logger.Error().Err(err).Uint("id", id).Msg("Failed to get banner")
		return nil, err
	}
	return bnr, nil
}

// GetAll retrieves all banners with pagination
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

	banners, total, err := uc.repo.GetAll(limit, offset)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get banners")
		return nil, errors.New("failed to get banners")
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return &GetAllResponse{
		Data:       banners,
		Total:      total,
		Limit:      limit,
		Offset:     offset,
		TotalPages: totalPages,
	}, nil
}

// Update updates an existing banner
func (uc *useCase) Update(id uint, req *UpdateRequest) (*bannerDomain.Banner, error) {
	bnr, err := uc.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if req.Title != "" {
		bnr.Title = req.Title
	}
	if req.MediaURL != "" {
		bnr.MediaURL = req.MediaURL
	}
	if req.Type != "" {
		bnr.Type = bannerDomain.BannerType(req.Type)
	}

	if err := uc.repo.Update(bnr); err != nil {
		logger.Error().Err(err).Uint("id", id).Msg("Failed to update banner")
		return nil, errors.New("failed to update banner")
	}

	logger.Info().
		Uint("id", bnr.ID).
		Msg("Banner updated successfully")

	return bnr, nil
}

// Delete deletes a banner
func (uc *useCase) Delete(id uint) error {
	if err := uc.repo.Delete(id); err != nil {
		logger.Error().Err(err).Uint("id", id).Msg("Failed to delete banner")
		return errors.New("failed to delete banner")
	}

	logger.Info().Uint("id", id).Msg("Banner deleted successfully")
	return nil
}

