package donationcategory

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	donationCategoryUsecase "github.com/madr/backend/internal/usecase/donationcategory"
	"github.com/madr/backend/pkg/logger"
)

// Handler handles HTTP requests for donation categories
type Handler struct {
	useCase donationCategoryUsecase.UseCase
}

// NewHandler creates a new donation category handler
func NewHandler(useCase donationCategoryUsecase.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

// Create handles POST /donation-categories
func (h *Handler) Create(c *gin.Context) {
	var req donationCategoryUsecase.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn().Err(err).Msg("Invalid create donation category request body")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	cat, err := h.useCase.Create(&req)
	if err != nil {
		if err.Error() == "category name already exists" {
			c.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})
			return
		}
		logger.Error().Err(err).Msg("Failed to create donation category")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create donation category",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Donation category created successfully",
		"data":    cat,
	})
}

// GetByID handles GET /donation-categories/:id
func (h *Handler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid category ID",
		})
		return
	}

	cat, err := h.useCase.GetByID(uint(id))
	if err != nil {
		if err.Error() == "donation category not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Donation category not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get donation category",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": cat,
	})
}

// GetAll handles GET /donation-categories
func (h *Handler) GetAll(c *gin.Context) {
	categories, err := h.useCase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get donation categories",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": categories,
	})
}

// Update handles PUT /donation-categories/:id
func (h *Handler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid category ID",
		})
		return
	}

	var req donationCategoryUsecase.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn().Err(err).Msg("Invalid update donation category request body")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	cat, err := h.useCase.Update(uint(id), &req)
	if err != nil {
		if err.Error() == "donation category not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Donation category not found",
			})
			return
		}
		if err.Error() == "category name already exists" {
			c.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update donation category",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Donation category updated successfully",
		"data":    cat,
	})
}

// Delete handles DELETE /donation-categories/:id
func (h *Handler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid category ID",
		})
		return
	}

	if err := h.useCase.Delete(uint(id)); err != nil {
		if err.Error() == "donation category not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Donation category not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete donation category",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Donation category deleted successfully",
	})
}

