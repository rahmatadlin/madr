package event

import (
	"errors"
	"time"

	eventDomain "github.com/madr/backend/internal/domain/event"
	eventRepo "github.com/madr/backend/internal/repository/event"
	"github.com/madr/backend/pkg/logger"
)

// UseCase defines the interface for event use case
type UseCase interface {
	Create(req *CreateRequest) (*eventDomain.Event, error)
	GetByID(id uint) (*eventDomain.Event, error)
	GetAll(limit, offset int) (*GetAllResponse, error)
	Update(id uint, req *UpdateRequest) (*eventDomain.Event, error)
	Delete(id uint) error
}

// CreateRequest represents the request to create an event
type CreateRequest struct {
	Title       string    `json:"title" binding:"required,min=3,max=255"`
	Description string    `json:"description"`
	Date        time.Time `json:"date" binding:"required"`
	Location    string    `json:"location" binding:"max=255"`
}

// UpdateRequest represents the request to update an event
type UpdateRequest struct {
	Title       string     `json:"title" binding:"min=3,max=255"`
	Description *string    `json:"description"`
	Date        *time.Time `json:"date"`
	Location    *string    `json:"location" binding:"max=255"`
}

// GetAllResponse represents the response for getting all events
type GetAllResponse struct {
	Data       []eventDomain.Event `json:"data"`
	Total      int64              `json:"total"`
	Limit      int                `json:"limit"`
	Offset     int                `json:"offset"`
	TotalPages int                `json:"total_pages"`
}

type useCase struct {
	repo eventRepo.Repository
}

// NewUseCase creates a new event use case
func NewUseCase(repo eventRepo.Repository) UseCase {
	return &useCase{
		repo: repo,
	}
}

// Create creates a new event
func (uc *useCase) Create(req *CreateRequest) (*eventDomain.Event, error) {
	evt := &eventDomain.Event{
		Title:       req.Title,
		Description: req.Description,
		Date:        req.Date,
		Location:    req.Location,
	}

	if err := uc.repo.Create(evt); err != nil {
		logger.Error().Err(err).Msg("Failed to create event")
		return nil, errors.New("failed to create event")
	}

	logger.Info().
		Uint("id", evt.ID).
		Str("title", evt.Title).
		Msg("Event created successfully")

	return evt, nil
}

// GetByID retrieves an event by ID
func (uc *useCase) GetByID(id uint) (*eventDomain.Event, error) {
	evt, err := uc.repo.GetByID(id)
	if err != nil {
		logger.Error().Err(err).Uint("id", id).Msg("Failed to get event")
		return nil, err
	}
	return evt, nil
}

// GetAll retrieves all events with pagination
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

	events, total, err := uc.repo.GetAll(limit, offset)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get events")
		return nil, errors.New("failed to get events")
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return &GetAllResponse{
		Data:       events,
		Total:      total,
		Limit:      limit,
		Offset:     offset,
		TotalPages: totalPages,
	}, nil
}

// Update updates an existing event
func (uc *useCase) Update(id uint, req *UpdateRequest) (*eventDomain.Event, error) {
	evt, err := uc.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if req.Title != "" {
		evt.Title = req.Title
	}
	if req.Description != nil {
		evt.Description = *req.Description
	}
	if req.Date != nil {
		evt.Date = *req.Date
	}
	if req.Location != nil {
		evt.Location = *req.Location
	}

	if err := uc.repo.Update(evt); err != nil {
		logger.Error().Err(err).Uint("id", id).Msg("Failed to update event")
		return nil, errors.New("failed to update event")
	}

	logger.Info().
		Uint("id", evt.ID).
		Msg("Event updated successfully")

	return evt, nil
}

// Delete deletes an event
func (uc *useCase) Delete(id uint) error {
	if err := uc.repo.Delete(id); err != nil {
		logger.Error().Err(err).Uint("id", id).Msg("Failed to delete event")
		return errors.New("failed to delete event")
	}

	logger.Info().Uint("id", id).Msg("Event deleted successfully")
	return nil
}

