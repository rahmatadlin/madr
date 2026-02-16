package youtube

import (
	"net/http"

	"github.com/gin-gonic/gin"
	youtubeService "github.com/madr/backend/internal/service/youtube"
	"github.com/madr/backend/pkg/logger"
)

// Handler handles YouTube API requests
type Handler struct {
	service youtubeService.Service
}

// NewHandler creates a new YouTube handler
func NewHandler() *Handler {
	return &Handler{
		service: youtubeService.NewService(),
	}
}

// GetKajianVideos handles GET /youtube/kajian
// Returns videos uploaded in the last 30 days
func (h *Handler) GetKajianVideos(c *gin.Context) {
	videos, err := h.service.GetRecentVideos(30)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get kajian videos")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch YouTube videos",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  videos,
		"total": len(videos),
	})
}
