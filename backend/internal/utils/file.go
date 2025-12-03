package utils

import (
	"crypto/rand"
	"fmt"
	"io"
	"mime"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/madr/backend/internal/config"
	"github.com/madr/backend/pkg/logger"
)

// ValidateFileType validates if the file MIME type is allowed
func ValidateFileType(mimeType string) bool {
	for _, allowedType := range config.AppConfig.Upload.AllowedTypes {
		if mimeType == allowedType {
			return true
		}
	}
	return false
}

// ValidateFileSize validates if the file size is within limit
func ValidateFileSize(size int64) bool {
	return size <= config.AppConfig.Upload.MaxSize
}

// GenerateUniqueFilename generates a unique filename using UUID and timestamp
func GenerateUniqueFilename(originalFilename string) string {
	// Get file extension
	ext := filepath.Ext(originalFilename)
	
	// Generate UUID
	id := uuid.New().String()
	
	// Get timestamp
	timestamp := time.Now().Unix()
	
	// Sanitize filename (remove any path components)
	baseName := filepath.Base(originalFilename)
	baseName = strings.TrimSuffix(baseName, ext)
	baseName = sanitizeFilename(baseName)
	
	// Combine: timestamp_uuid_sanitizedname.ext
	filename := fmt.Sprintf("%d_%s_%s%s", timestamp, id, baseName, ext)
	
	return filename
}

// sanitizeFilename removes dangerous characters from filename
func sanitizeFilename(filename string) string {
	// Remove path separators and dangerous characters
	filename = strings.ReplaceAll(filename, "/", "_")
	filename = strings.ReplaceAll(filename, "\\", "_")
	filename = strings.ReplaceAll(filename, "..", "_")
	filename = strings.ReplaceAll(filename, " ", "_")
	
	// Remove any non-alphanumeric characters except underscore and dash
	var sanitized strings.Builder
	for _, char := range filename {
		if (char >= 'a' && char <= 'z') || 
		   (char >= 'A' && char <= 'Z') || 
		   (char >= '0' && char <= '9') || 
		   char == '_' || char == '-' {
			sanitized.WriteRune(char)
		}
	}
	
	result := sanitized.String()
	if result == "" {
		// If sanitization removed everything, use random string
		result = generateRandomString(8)
	}
	
	return result
}

// generateRandomString generates a random alphanumeric string
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	rand.Read(b)
	for i := range b {
		b[i] = charset[b[i]%byte(len(charset))]
	}
	return string(b)
}

// SaveFile saves uploaded file to disk
func SaveFile(src io.Reader, filename string) (string, error) {
	// Ensure upload directory exists
	uploadPath := config.AppConfig.Upload.UploadPath
	if err := os.MkdirAll(uploadPath, 0755); err != nil {
		logger.Error().Err(err).Str("path", uploadPath).Msg("Failed to create upload directory")
		return "", fmt.Errorf("failed to create upload directory: %w", err)
	}

	// Create full file path
	fullPath := filepath.Join(uploadPath, filename)

	// Create file
	dst, err := os.Create(fullPath)
	if err != nil {
		logger.Error().Err(err).Str("path", fullPath).Msg("Failed to create file")
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer dst.Close()

	// Copy file content
	if _, err := io.Copy(dst, src); err != nil {
		logger.Error().Err(err).Str("path", fullPath).Msg("Failed to save file")
		return "", fmt.Errorf("failed to save file: %w", err)
	}

	logger.Info().
		Str("filename", filename).
		Str("path", fullPath).
		Msg("File saved successfully")

	return fullPath, nil
}

// GetMIMEType detects MIME type from file extension
func GetMIMEType(filename string) string {
	ext := filepath.Ext(filename)
	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		// Fallback to common types
		switch strings.ToLower(ext) {
		case ".jpg", ".jpeg":
			return "image/jpeg"
		case ".png":
			return "image/png"
		case ".webp":
			return "image/webp"
		case ".mp4":
			return "video/mp4"
		default:
			return "application/octet-stream"
		}
	}
	return mimeType
}

// GetPublicURL returns the public URL for a file
func GetPublicURL(filename string) string {
	publicURL := strings.TrimSuffix(config.AppConfig.Upload.PublicURL, "/")
	return fmt.Sprintf("%s/%s", publicURL, filename)
}

// DeleteFile deletes a file from disk
func DeleteFile(filename string) error {
	uploadPath := config.AppConfig.Upload.UploadPath
	fullPath := filepath.Join(uploadPath, filename)

	if err := os.Remove(fullPath); err != nil {
		if !os.IsNotExist(err) {
			logger.Error().Err(err).Str("path", fullPath).Msg("Failed to delete file")
			return fmt.Errorf("failed to delete file: %w", err)
		}
		// File doesn't exist, consider it deleted
		logger.Warn().Str("path", fullPath).Msg("File not found for deletion")
	}

	logger.Info().Str("filename", filename).Msg("File deleted successfully")
	return nil
}

