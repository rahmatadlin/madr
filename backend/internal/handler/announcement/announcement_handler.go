package announcement

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/madr/backend/internal/usecase/announcement"
	"github.com/madr/backend/pkg/logger"
)

// Handler handles HTTP requests for announcements
type Handler struct {
	useCase announcement.UseCase
}

// NewHandler creates a new announcement handler
func NewHandler(useCase announcement.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

// Create handles POST /announcements
func (h *Handler) Create(c *gin.Context) {
	var req announcement.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn().Err(err).Msg("Invalid request body")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	ann, err := h.useCase.Create(&req)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to create announcement")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create announcement",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Announcement created successfully",
		"data":    ann,
	})
}

// GetByID handles GET /announcements/:id
func (h *Handler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid announcement ID",
		})
		return
	}

	ann, err := h.useCase.GetByID(uint(id))
	if err != nil {
		if err.Error() == "announcement not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Announcement not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get announcement",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ann,
	})
}

// GetAll handles GET /announcements (admin only, shows all including unpublished)
func (h *Handler) GetAll(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	response, err := h.useCase.GetAll(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get announcements",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetPublished handles GET /announcements/published (public endpoint)
func (h *Handler) GetPublished(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	response, err := h.useCase.GetPublished(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get published announcements",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// Update handles PUT /announcements/:id
func (h *Handler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid announcement ID",
		})
		return
	}

	var req announcement.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn().Err(err).Msg("Invalid request body")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	ann, err := h.useCase.Update(uint(id), &req)
	if err != nil {
		if err.Error() == "announcement not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Announcement not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update announcement",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Announcement updated successfully",
		"data":    ann,
	})
}

// Delete handles DELETE /announcements/:id
func (h *Handler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid announcement ID",
		})
		return
	}

	if err := h.useCase.Delete(uint(id)); err != nil {
		if err.Error() == "announcement not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Announcement not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete announcement",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Announcement deleted successfully",
	})
}

