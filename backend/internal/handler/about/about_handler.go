package about

import (
	"net/http"

	"github.com/gin-gonic/gin"
	aboutUsecase "github.com/madr/backend/internal/usecase/about"
	"github.com/madr/backend/pkg/logger"
	"gorm.io/gorm"
)

// Handler handles HTTP requests for about content.
type Handler struct {
	useCase aboutUsecase.UseCase
}

// NewHandler creates a new about handler.
func NewHandler(useCase aboutUsecase.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

// Get handles GET /about
func (h *Handler) Get(c *gin.Context) {
	abt, err := h.useCase.Get()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "About content not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch about content",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": abt,
	})
}

// Update handles PUT /admin/about
func (h *Handler) Update(c *gin.Context) {
	var req aboutUsecase.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn().Err(err).Msg("Invalid update about request body")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	abt, err := h.useCase.Update(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update about content",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "About content saved",
		"data":    abt,
	})
}
