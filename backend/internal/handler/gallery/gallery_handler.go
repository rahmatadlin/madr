package gallery

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	galleryUsecase "github.com/madr/backend/internal/usecase/gallery"
	"github.com/madr/backend/internal/utils"
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

// Create handles POST /gallery (with file upload)
func (h *Handler) Create(c *gin.Context) {
	// Get file from form
	file, err := c.FormFile("file")
	if err != nil {
		logger.Warn().Err(err).Msg("Failed to get file from form")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "File is required",
		})
		return
	}

	// Get title from form
	title := c.PostForm("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Title is required",
		})
		return
	}

	// Open file
	src, err := file.Open()
	if err != nil {
		logger.Error().Err(err).Str("filename", file.Filename).Msg("Failed to open file")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to process file",
		})
		return
	}
	defer src.Close()

	// Validate file size
	if !utils.ValidateFileSize(file.Size) {
		logger.Warn().
			Str("filename", file.Filename).
			Int64("size", file.Size).
			Msg("File size exceeds limit")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("File size exceeds maximum allowed size (%d bytes)", file.Size),
		})
		return
	}

	// Detect MIME type
	mimeType := utils.GetMIMEType(file.Filename)
	
	// Validate MIME type (only images allowed for gallery)
	allowedImageTypes := []string{"image/jpeg", "image/jpg", "image/png", "image/webp"}
	isImage := false
	for _, allowedType := range allowedImageTypes {
		if mimeType == allowedType {
			isImage = true
			break
		}
	}

	if !isImage {
		logger.Warn().
			Str("filename", file.Filename).
			Str("mime_type", mimeType).
			Msg("Invalid file type for gallery (only images allowed)")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Only image files (jpg, jpeg, png, webp) are allowed for gallery",
		})
		return
	}

	// Generate unique filename
	uniqueFilename := utils.GenerateUniqueFilename(file.Filename)

	// Save file
	_, err = utils.SaveFile(src, uniqueFilename)
	if err != nil {
		logger.Error().Err(err).Str("filename", file.Filename).Msg("Failed to save file")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save file",
		})
		return
	}

	// Get public URL
	publicURL := utils.GetPublicURL(uniqueFilename)

	// Create gallery item
	req := galleryUsecase.CreateRequest{
		Title:    title,
		ImageURL: uniqueFilename, // Store filename, not full URL
	}

	gal, err := h.useCase.Create(&req)
	if err != nil {
		// If creation fails, delete the uploaded file
		utils.DeleteFile(uniqueFilename)
		logger.Error().Err(err).Msg("Failed to create gallery item")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create gallery item",
		})
		return
	}

	// Update response with full URL
	gal.ImageURL = publicURL

	logger.Info().
		Uint("id", gal.ID).
		Str("title", gal.Title).
		Str("filename", uniqueFilename).
		Msg("Gallery item created successfully with file upload")

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

