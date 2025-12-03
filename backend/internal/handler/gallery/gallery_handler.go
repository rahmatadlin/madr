package gallery

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	galleryUsecase "github.com/madr/backend/internal/usecase/gallery"
	"github.com/madr/backend/pkg/logger"
)

// Handler handles HTTP requests for gallery
type Handler struct {
	useCase galleryUsecase.UseCase
}

// NewHandler creates a new gallery handler
func NewHandler(useCase galleryUsecase.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

// Create handles POST /gallery
func (h *Handler) Create(c *gin.Context) {
	var req galleryUsecase.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn().Err(err).Msg("Invalid create gallery request body")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	gal, err := h.useCase.Create(&req)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to create gallery item")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create gallery item",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Gallery item created successfully",
		"data":    gal,
	})
}

// GetAll handles GET /gallery
func (h *Handler) GetAll(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	response, err := h.useCase.GetAll(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get gallery items",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// Delete handles DELETE /gallery/:id
func (h *Handler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid gallery ID",
		})
		return
	}

	if err := h.useCase.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete gallery item",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Gallery item deleted successfully",
	})
}

