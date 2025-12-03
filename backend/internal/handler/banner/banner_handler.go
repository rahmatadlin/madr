package banner

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	bannerUsecase "github.com/madr/backend/internal/usecase/banner"
	"github.com/madr/backend/internal/utils"
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

// Create handles POST /banners (with file upload)
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

	// Get title and type from form
	title := c.PostForm("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Title is required",
		})
		return
	}

	bannerType := strings.ToLower(c.PostForm("type"))
	if bannerType == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Type is required (image or video)",
		})
		return
	}

	if bannerType != "image" && bannerType != "video" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Type must be either 'image' or 'video'",
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
	
	// Validate MIME type based on banner type
	allowedImageTypes := []string{"image/jpeg", "image/jpg", "image/png", "image/webp"}
	allowedVideoTypes := []string{"video/mp4"}

	isValidType := false
	if bannerType == "image" {
		for _, allowedType := range allowedImageTypes {
			if mimeType == allowedType {
				isValidType = true
				break
			}
		}
		if !isValidType {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid file type for image banner. Only jpg, jpeg, png, webp are allowed",
			})
			return
		}
	} else if bannerType == "video" {
		for _, allowedType := range allowedVideoTypes {
			if mimeType == allowedType {
				isValidType = true
				break
			}
		}
		if !isValidType {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid file type for video banner. Only mp4 is allowed",
			})
			return
		}
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

	// Create banner
	req := bannerUsecase.CreateRequest{
		Title:    title,
		MediaURL: uniqueFilename, // Store filename, not full URL
		Type:     bannerType,
	}

	bnr, err := h.useCase.Create(&req)
	if err != nil {
		// If creation fails, delete the uploaded file
		utils.DeleteFile(uniqueFilename)
		logger.Error().Err(err).Msg("Failed to create banner")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create banner",
		})
		return
	}

	// Update response with full URL
	bnr.MediaURL = publicURL

	logger.Info().
		Uint("id", bnr.ID).
		Str("title", bnr.Title).
		Str("type", string(bnr.Type)).
		Str("filename", uniqueFilename).
		Msg("Banner created successfully with file upload")

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

// Update handles PUT /banners/:id (with optional file upload)
func (h *Handler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid banner ID",
		})
		return
	}

	// Check if file is provided
	file, err := c.FormFile("file")
	fileProvided := err == nil
	var uniqueFilename string
	var publicURL string

	if fileProvided {
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
		
		// Get banner type from form or existing banner
		bannerType := strings.ToLower(c.PostForm("type"))
		if bannerType == "" {
			// Get existing banner to check type
			existingBanner, err := h.useCase.GetByID(uint(id))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Type is required when uploading new file",
				})
				return
			}
			bannerType = string(existingBanner.Type)
		}

		// Validate MIME type based on banner type
		allowedImageTypes := []string{"image/jpeg", "image/jpg", "image/png", "image/webp"}
		allowedVideoTypes := []string{"video/mp4"}

		isValidType := false
		if bannerType == "image" {
			for _, allowedType := range allowedImageTypes {
				if mimeType == allowedType {
					isValidType = true
					break
				}
			}
			if !isValidType {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Invalid file type for image banner. Only jpg, jpeg, png, webp are allowed",
				})
				return
			}
		} else if bannerType == "video" {
			for _, allowedType := range allowedVideoTypes {
				if mimeType == allowedType {
					isValidType = true
					break
				}
			}
			if !isValidType {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Invalid file type for video banner. Only mp4 is allowed",
				})
				return
			}
		}

		// Generate unique filename
		uniqueFilename = utils.GenerateUniqueFilename(file.Filename)

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
		publicURL = utils.GetPublicURL(uniqueFilename)
	}

	// Build update request
	req := bannerUsecase.UpdateRequest{}
	
	title := c.PostForm("title")
	if title != "" {
		req.Title = title
	}

	bannerType := strings.ToLower(c.PostForm("type"))
	if bannerType != "" {
		req.Type = bannerType
	}

	if fileProvided {
		req.MediaURL = uniqueFilename
	}

	bnr, err := h.useCase.Update(uint(id), &req)
	if err != nil {
		// If update fails and file was uploaded, delete the uploaded file
		if fileProvided {
			utils.DeleteFile(uniqueFilename)
		}
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

	// Update response with full URL if file was uploaded
	if fileProvided {
		bnr.MediaURL = publicURL
	}

	logger.Info().
		Uint("id", bnr.ID).
		Bool("file_uploaded", fileProvided).
		Msg("Banner updated successfully")

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

