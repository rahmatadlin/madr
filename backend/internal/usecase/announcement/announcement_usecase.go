package announcement

import (
	"errors"

	announcementDomain "github.com/madr/backend/internal/domain/announcement"
	announcementRepo "github.com/madr/backend/internal/repository/announcement"
	"github.com/madr/backend/pkg/logger"
)

// UseCase defines the interface for announcement use case
type UseCase interface {
	Create(req *CreateRequest) (*announcementDomain.Announcement, error)
	GetByID(id uint) (*announcementDomain.Announcement, error)
	GetAll(limit, offset int) (*GetAllResponse, error)
	GetPublished(limit, offset int) (*GetAllResponse, error)
	Update(id uint, req *UpdateRequest) (*announcementDomain.Announcement, error)
	Delete(id uint) error
}

// CreateRequest represents the request to create an announcement
type CreateRequest struct {
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	IsPublished bool   `json:"is_published"`
	Author      string `json:"author"`
}

// UpdateRequest represents the request to update an announcement
type UpdateRequest struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	IsPublished *bool  `json:"is_published"`
	Author      string `json:"author"`
}

// GetAllResponse represents the response for getting all announcements
type GetAllResponse struct {
	Data       []announcementDomain.Announcement `json:"data"`
	Total      int64                             `json:"total"`
	Limit      int                               `json:"limit"`
	Offset     int                               `json:"offset"`
	TotalPages int                               `json:"total_pages"`
}

type useCase struct {
	repo announcementRepo.Repository
}

// NewUseCase creates a new announcement use case
func NewUseCase(repo announcementRepo.Repository) UseCase {
	return &useCase{
		repo: repo,
	}
}

// Create creates a new announcement
func (uc *useCase) Create(req *CreateRequest) (*announcementDomain.Announcement, error) {
	ann := &announcementDomain.Announcement{
		Title:       req.Title,
		Content:     req.Content,
		IsPublished: req.IsPublished,
		Author:      req.Author,
	}

	if err := uc.repo.Create(ann); err != nil {
		logger.Error().Err(err).Msg("Failed to create announcement")
		return nil, errors.New("failed to create announcement")
	}

	logger.Info().
		Uint("id", ann.ID).
		Str("title", ann.Title).
		Msg("Announcement created successfully")

	return ann, nil
}

// GetByID retrieves an announcement by ID
func (uc *useCase) GetByID(id uint) (*announcementDomain.Announcement, error) {
	ann, err := uc.repo.GetByID(id)
	if err != nil {
		logger.Error().Err(err).Uint("id", id).Msg("Failed to get announcement")
		return nil, err
	}
	return ann, nil
}

// GetAll retrieves all announcements with pagination
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

	announcements, total, err := uc.repo.GetAll(limit, offset)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get announcements")
		return nil, errors.New("failed to get announcements")
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return &GetAllResponse{
		Data:       announcements,
		Total:      total,
		Limit:      limit,
		Offset:     offset,
		TotalPages: totalPages,
	}, nil
}

// GetPublished retrieves only published announcements with pagination
func (uc *useCase) GetPublished(limit, offset int) (*GetAllResponse, error) {
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

	announcements, total, err := uc.repo.GetPublished(limit, offset)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get published announcements")
		return nil, errors.New("failed to get published announcements")
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return &GetAllResponse{
		Data:       announcements,
		Total:      total,
		Limit:      limit,
		Offset:     offset,
		TotalPages: totalPages,
	}, nil
}

// Update updates an existing announcement
func (uc *useCase) Update(id uint, req *UpdateRequest) (*announcementDomain.Announcement, error) {
	ann, err := uc.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if req.Title != "" {
		ann.Title = req.Title
	}
	if req.Content != "" {
		ann.Content = req.Content
	}
	if req.IsPublished != nil {
		ann.IsPublished = *req.IsPublished
	}
	if req.Author != "" {
		ann.Author = req.Author
	}

	if err := uc.repo.Update(ann); err != nil {
		logger.Error().Err(err).Uint("id", id).Msg("Failed to update announcement")
		return nil, errors.New("failed to update announcement")
	}

	logger.Info().
		Uint("id", ann.ID).
		Msg("Announcement updated successfully")

	return ann, nil
}

// Delete deletes an announcement
func (uc *useCase) Delete(id uint) error {
	if err := uc.repo.Delete(id); err != nil {
		logger.Error().Err(err).Uint("id", id).Msg("Failed to delete announcement")
		return errors.New("failed to delete announcement")
	}

	logger.Info().Uint("id", id).Msg("Announcement deleted successfully")
	return nil
}

