package upload

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/madr/backend/internal/utils"
	"github.com/madr/backend/pkg/logger"
)

// Handler handles file upload requests
type Handler struct{}

// NewHandler creates a new upload handler
func NewHandler() *Handler {
	return &Handler{}
}

// UploadFile handles POST /upload
func (h *Handler) UploadFile(c *gin.Context) {
	// Get file from form
	file, err := c.FormFile("file")
	if err != nil {
		logger.Warn().Err(err).Msg("Failed to get file from form")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "File is required",
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
	
	// Validate MIME type
	if !utils.ValidateFileType(mimeType) {
		logger.Warn().
			Str("filename", file.Filename).
			Str("mime_type", mimeType).
			Msg("Invalid file type")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("File type '%s' is not allowed", mimeType),
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

	logger.Info().
		Str("original_filename", file.Filename).
		Str("saved_filename", uniqueFilename).
		Str("mime_type", mimeType).
		Int64("size", file.Size).
		Msg("File uploaded successfully")

	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully",
		"data": gin.H{
			"filename":    uniqueFilename,
			"original_name": file.Filename,
			"url":         publicURL,
			"mime_type":   mimeType,
			"size":        file.Size,
		},
	})
}

