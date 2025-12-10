package about

import (
	"errors"

	aboutDomain "github.com/madr/backend/internal/domain/about"
	aboutRepo "github.com/madr/backend/internal/repository/about"
	"github.com/madr/backend/pkg/logger"
	"gorm.io/gorm"
)

// UseCase defines business logic for about section.
type UseCase interface {
	Get() (*aboutDomain.About, error)
	Update(req *UpdateRequest) (*aboutDomain.About, error)
}

// UpdateRequest represents payload to update about content.
type UpdateRequest struct {
	Title                 *string `json:"title"`
	Subtitle              *string `json:"subtitle"`
	Description           *string `json:"description"`
	AdditionalDescription *string `json:"additional_description"`
	ImageURL              *string `json:"image_url"`
	YearsActive           *int    `json:"years_active"`
	ActiveMembers         *int    `json:"active_members"`
}

type useCase struct {
	repo aboutRepo.Repository
}

// NewUseCase constructs the about use case.
func NewUseCase(repo aboutRepo.Repository) UseCase {
	return &useCase{repo: repo}
}

// Get returns the latest about content.
func (uc *useCase) Get() (*aboutDomain.About, error) {
	abt, err := uc.repo.GetLatest()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		logger.Error().Err(err).Msg("Failed to get about content")
		return nil, errors.New("failed to get about content")
	}
	return abt, nil
}

// Update upserts the about content (single row).
func (uc *useCase) Update(req *UpdateRequest) (*aboutDomain.About, error) {
	var abt *aboutDomain.About

	existing, err := uc.repo.GetLatest()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Error().Err(err).Msg("Failed to fetch existing about content")
		return nil, errors.New("failed to update about content")
	}

	if existing != nil {
		abt = existing
	} else {
		abt = &aboutDomain.About{}
	}

	// Apply updates only when provided
	if req.Title != nil {
		abt.Title = *req.Title
	}
	if req.Subtitle != nil {
		abt.Subtitle = *req.Subtitle
	}
	if req.Description != nil {
		abt.Description = *req.Description
	}
	if req.AdditionalDescription != nil {
		abt.AdditionalDescription = *req.AdditionalDescription
	}
	if req.ImageURL != nil {
		abt.ImageURL = *req.ImageURL
	}
	if req.YearsActive != nil {
		abt.YearsActive = *req.YearsActive
	}
	if req.ActiveMembers != nil {
		abt.ActiveMembers = *req.ActiveMembers
	}

	// Persist
	if existing == nil {
		if err := uc.repo.Create(abt); err != nil {
			logger.Error().Err(err).Msg("Failed to create about content")
			return nil, errors.New("failed to update about content")
		}
		logger.Info().Uint("id", abt.ID).Msg("About content created")
	} else {
		if err := uc.repo.Update(abt); err != nil {
			logger.Error().Err(err).Uint("id", abt.ID).Msg("Failed to update about content")
			return nil, errors.New("failed to update about content")
		}
		logger.Info().Uint("id", abt.ID).Msg("About content updated")
	}

	return abt, nil
}
