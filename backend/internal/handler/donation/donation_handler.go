package donation

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	donationUsecase "github.com/madr/backend/internal/usecase/donation"
	donationDomain "github.com/madr/backend/internal/domain/donation"
	"github.com/madr/backend/pkg/logger"
)

// Handler handles HTTP requests for donations
type Handler struct {
	useCase donationUsecase.UseCase
}

// NewHandler creates a new donation handler
func NewHandler(useCase donationUsecase.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

// Create handles POST /donations
func (h *Handler) Create(c *gin.Context) {
	var req donationUsecase.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn().Err(err).Msg("Invalid create donation request body")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	don, err := h.useCase.Create(&req)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to create donation")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create donation",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Donation created successfully",
		"data":    don,
	})
}

// GetByID handles GET /donations/:id
func (h *Handler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid donation ID",
		})
		return
	}

	don, err := h.useCase.GetByID(uint(id))
	if err != nil {
		if err.Error() == "donation not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Donation not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get donation",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": don,
	})
}

// GetAll handles GET /donations
func (h *Handler) GetAll(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	
	// Optional status filter
	statusStr := c.Query("status")
	var status *donationDomain.PaymentStatus
	if statusStr != "" {
		ps := donationDomain.PaymentStatus(statusStr)
		if ps == donationDomain.PaymentStatusPending || 
		   ps == donationDomain.PaymentStatusSuccess || 
		   ps == donationDomain.PaymentStatusFailed {
			status = &ps
		}
	}

	response, err := h.useCase.GetAll(limit, offset, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get donations",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// Update handles PUT /donations/:id
func (h *Handler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid donation ID",
		})
		return
	}

	var req donationUsecase.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn().Err(err).Msg("Invalid update donation request body")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	don, err := h.useCase.Update(uint(id), &req)
	if err != nil {
		if err.Error() == "donation not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Donation not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update donation",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Donation updated successfully",
		"data":    don,
	})
}

// Delete handles DELETE /donations/:id
func (h *Handler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid donation ID",
		})
		return
	}

	if err := h.useCase.Delete(uint(id)); err != nil {
		if err.Error() == "donation not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Donation not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete donation",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Donation deleted successfully",
	})
}

// GetSummary handles GET /donations/summary (Public endpoint)
func (h *Handler) GetSummary(c *gin.Context) {
	summary, err := h.useCase.GetSummary()
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get donation summary")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get donation summary",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": summary,
	})
}

