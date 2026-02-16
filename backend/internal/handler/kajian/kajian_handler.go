package kajian

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	kajianUsecase "github.com/madr/backend/internal/usecase/kajian"
	"github.com/madr/backend/pkg/logger"
)

// Handler handles HTTP requests for kajian
type Handler struct {
	useCase kajianUsecase.UseCase
}

// NewHandler creates a new kajian handler
func NewHandler(useCase kajianUsecase.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

// SyncFromYouTube handles POST /admin/kajian/sync
func (h *Handler) SyncFromYouTube(c *gin.Context) {
	days := 30
	if d := c.Query("days"); d != "" {
		if parsed, err := strconv.Atoi(d); err == nil && parsed > 0 {
			days = parsed
		}
	}

	synced, err := h.useCase.SyncFromYouTube(days)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to sync kajian from YouTube")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to sync from YouTube",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Kajian synced successfully",
		"synced":  synced,
	})
}

// GetByID handles GET /kajian/:id
func (h *Handler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid kajian ID"})
		return
	}
	k, err := h.useCase.GetByID(uint(id))
	if err != nil {
		if err.Error() == "kajian not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Kajian not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get kajian"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": k})
}

// GetAll handles GET /kajian
func (h *Handler) GetAll(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	resp, err := h.useCase.GetAll(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get kajian"})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// Delete handles DELETE /admin/kajian/:id
func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid kajian ID"})
		return
	}
	if err := h.useCase.Delete(uint(id)); err != nil {
		if err.Error() == "kajian not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Kajian not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete kajian"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Kajian deleted successfully"})
}
