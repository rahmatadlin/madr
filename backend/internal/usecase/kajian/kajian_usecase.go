package kajian

import (
	"errors"
	"fmt"

	kajianDomain "github.com/madr/backend/internal/domain/kajian"
	kajianRepo "github.com/madr/backend/internal/repository/kajian"
	youtubeService "github.com/madr/backend/internal/service/youtube"
	"github.com/madr/backend/pkg/logger"
)

// UseCase defines the interface for kajian use case
type UseCase interface {
	SyncFromYouTube(days int) (int, error)
	GetAll(limit, offset int) (*GetAllResponse, error)
	GetByID(id uint) (*kajianDomain.Kajian, error)
	Delete(id uint) error
}

// GetAllResponse represents the response for getting all kajian
type GetAllResponse struct {
	Data       []kajianDomain.Kajian `json:"data"`
	Total      int64                 `json:"total"`
	Limit      int                   `json:"limit"`
	Offset     int                   `json:"offset"`
	TotalPages int                   `json:"total_pages"`
}

type useCase struct {
	repo           kajianRepo.Repository
	youtubeService youtubeService.Service
}

// NewUseCase creates a new kajian use case
func NewUseCase(repo kajianRepo.Repository, ytService youtubeService.Service) UseCase {
	return &useCase{
		repo:           repo,
		youtubeService: ytService,
	}
}

// SyncFromYouTube fetches videos from YouTube API and saves them to database
func (uc *useCase) SyncFromYouTube(days int) (int, error) {
	videos, err := uc.youtubeService.GetRecentVideos(days)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch YouTube videos: %w", err)
	}

	synced := 0
	for _, v := range videos {
		k := &kajianDomain.Kajian{
			VideoID:      v.VideoID,
			Title:        v.Title,
			Description:  v.Description,
			PublishedAt:  v.PublishedAt,
			ThumbnailURL: v.ThumbnailURL,
			YoutubeURL:   fmt.Sprintf("https://www.youtube.com/watch?v=%s", v.VideoID),
			ChannelTitle: v.ChannelTitle,
		}
		if err := uc.repo.CreateOrUpdate(k); err != nil {
			logger.Warn().Err(err).Str("video_id", v.VideoID).Msg("Failed to save kajian")
			continue
		}
		synced++
	}

	logger.Info().Int("synced", synced).Int("total", len(videos)).Msg("Synced kajian from YouTube")
	return synced, nil
}

// GetAll retrieves all kajian with pagination
func (uc *useCase) GetAll(limit, offset int) (*GetAllResponse, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}
	list, total, err := uc.repo.GetAll(limit, offset)
	if err != nil {
		return nil, errors.New("failed to get kajian")
	}
	totalPages := int((total + int64(limit) - 1) / int64(limit))
	return &GetAllResponse{
		Data:       list,
		Total:      total,
		Limit:      limit,
		Offset:     offset,
		TotalPages: totalPages,
	}, nil
}

// GetByID retrieves a kajian by ID
func (uc *useCase) GetByID(id uint) (*kajianDomain.Kajian, error) {
	return uc.repo.GetByID(id)
}

// Delete deletes a kajian
func (uc *useCase) Delete(id uint) error {
	if err := uc.repo.Delete(id); err != nil {
		return errors.New("failed to delete kajian")
	}
	return nil
}
