package banner

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	bannerUsecase "github.com/madr/backend/internal/usecase/banner"
	"github.com/madr/backend/pkg/logger"
)

// Handler handles HTTP requests for banners
type Handler struct {
	useCase bannerUsecase.UseCase
}

// NewHandler creates a new banner handler
func NewHandler(useCase bannerUsecase.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

// Create handles POST /banners
func (h *Handler) Create(c *gin.Context) {
	var req bannerUsecase.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn().Err(err).Msg("Invalid create banner request body")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	bnr, err := h.useCase.Create(&req)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to create banner")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create banner",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Banner created successfully",
		"data":    bnr,
	})
}

// GetByID handles GET /banners/:id
func (h *Handler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid banner ID",
		})
		return
	}

	bnr, err := h.useCase.GetByID(uint(id))
	if err != nil {
		if err.Error() == "banner not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Banner not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get banner",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": bnr,
	})
}

// GetAll handles GET /banners
func (h *Handler) GetAll(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	response, err := h.useCase.GetAll(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get banners",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// Update handles PUT /banners/:id
func (h *Handler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid banner ID",
		})
		return
	}

	var req bannerUsecase.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn().Err(err).Msg("Invalid update banner request body")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	bnr, err := h.useCase.Update(uint(id), &req)
	if err != nil {
		if err.Error() == "banner not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Banner not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update banner",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Banner updated successfully",
		"data":    bnr,
	})
}

// Delete handles DELETE /banners/:id
func (h *Handler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid banner ID",
		})
		return
	}

	if err := h.useCase.Delete(uint(id)); err != nil {
		if err.Error() == "banner not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Banner not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete banner",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Banner deleted successfully",
	})
}

