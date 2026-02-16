package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	Server    ServerConfig
	Database  DatabaseConfig
	JWT       JWTConfig
	CORS      CORSConfig
	RateLimit RateLimitConfig
	Logging   LoggingConfig
	Upload    UploadConfig
	YouTube   YouTubeConfig
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Host string
	Port string
	Mode string
}

// DatabaseConfig holds database-related configuration
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

// JWTConfig holds JWT-related configuration
type JWTConfig struct {
	Secret        string
	AccessExpiry  time.Duration
	RefreshExpiry time.Duration
}

// CORSConfig holds CORS-related configuration
type CORSConfig struct {
	AllowedOrigins []string
	AllowedMethods []string
	AllowedHeaders []string
}

// RateLimitConfig holds rate limiting configuration
type RateLimitConfig struct {
	Enabled  bool
	Requests int
	Window   time.Duration
}

// LoggingConfig holds logging configuration
type LoggingConfig struct {
	Level  string
	Format string
}

// UploadConfig holds file upload configuration
type UploadConfig struct {
	MaxSize      int64
	AllowedTypes []string
	UploadPath   string
	PublicURL    string
}

// YouTubeConfig holds YouTube API configuration
type YouTubeConfig struct {
	APIKey    string
	ChannelID string
	APIURL    string
}

var AppConfig *Config

// Load loads configuration from environment variables
func Load() error {
	// Try to load .env file from multiple possible locations
	// First try current directory
	envPaths := []string{
		".env",
		"./.env",
		"../.env",
		"../../.env",
	}
	
	var envLoaded bool
	for _, path := range envPaths {
		if err := godotenv.Load(path); err == nil {
			envLoaded = true
			break
		}
	}
	
	// Also try to find .env relative to the executable
	if !envLoaded {
		// Try to get the directory where the binary is running
		execPath, err := os.Executable()
		if err == nil {
			execDir := filepath.Dir(execPath)
			envPath := filepath.Join(execDir, ".env")
			if err := godotenv.Load(envPath); err == nil {
				envLoaded = true
			}
		}
	}
	
	// If still not loaded, try without path (godotenv will search)
	if !envLoaded {
		_ = godotenv.Load() // Ignore error, .env is optional
	}

	AppConfig = &Config{
		Server: ServerConfig{
			Host: getEnv("SERVER_HOST", "localhost"),
			Port: getEnv("SERVER_PORT", "8080"),
			Mode: getEnv("SERVER_MODE", "debug"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			Name:     getEnv("DB_NAME", "madr_db"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			Secret:        getEnv("JWT_SECRET", "your-secret-key-change-in-production"),
			AccessExpiry:  parseDuration(getEnv("JWT_ACCESS_EXPIRY", "15m")),
			RefreshExpiry: parseDuration(getEnv("JWT_REFRESH_EXPIRY", "7d")),
		},
		CORS: CORSConfig{
			AllowedOrigins: getEnvSlice("CORS_ALLOWED_ORIGINS", []string{"http://localhost:3000","http://localhost:3001"}),
			AllowedMethods: getEnvSlice("CORS_ALLOWED_METHODS", []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"}),
			AllowedHeaders: getEnvSlice("CORS_ALLOWED_HEADERS", []string{"Content-Type", "Authorization"}),
		},
		RateLimit: RateLimitConfig{
			Enabled:  getEnvBool("RATE_LIMIT_ENABLED", true),
			Requests: getEnvInt("RATE_LIMIT_REQUESTS", 100),
			Window:   parseDuration(getEnv("RATE_LIMIT_WINDOW", "1m")),
		},
		Logging: LoggingConfig{
			Level:  getEnv("LOG_LEVEL", "info"),
			Format: getEnv("LOG_FORMAT", "json"),
		},
		Upload: UploadConfig{
			MaxSize:      int64(getEnvInt("UPLOAD_MAX_SIZE", 50*1024*1024)), // Default 50MB
			AllowedTypes: getEnvSlice("UPLOAD_ALLOWED_TYPES", []string{"image/jpeg", "image/jpg", "image/png", "image/webp", "video/mp4"}),
			UploadPath:   getEnv("UPLOAD_PATH", "./uploads"),
			PublicURL:    getEnv("UPLOAD_PUBLIC_URL", "http://localhost:8080/uploads"),
		},
		YouTube: YouTubeConfig{
			APIKey:    getEnv("YOUTUBE_API_KEY", ""),
			ChannelID: getEnv("YOUTUBE_CHANNEL_ID", ""),
			APIURL:    getEnv("YOUTUBE_API_URL", "https://www.googleapis.com/youtube/v3/search"),
		},
	}

	// Fallback: Try to read directly from environment if not loaded from .env
	if AppConfig.YouTube.APIKey == "" {
		if envKey := os.Getenv("YOUTUBE_API_KEY"); envKey != "" {
			AppConfig.YouTube.APIKey = envKey
		}
	}
	if AppConfig.YouTube.ChannelID == "" {
		if envChannelID := os.Getenv("YOUTUBE_CHANNEL_ID"); envChannelID != "" {
			AppConfig.YouTube.ChannelID = envChannelID
		}
	}

	return nil
}

// GetDSN returns the database connection string
func (c *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.Name, c.SSLMode,
	)
}

// GetServerAddress returns the server address
func (c *ServerConfig) GetServerAddress() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

// Helper functions
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}

func getEnvSlice(key string, defaultValue []string) []string {
	if value := os.Getenv(key); value != "" {
		// Simple comma-separated parsing
		var result []string
		start := 0
		for i, char := range value {
			if char == ',' {
				if i > start {
					result = append(result, value[start:i])
				}
				start = i + 1
			}
		}
		if start < len(value) {
			result = append(result, value[start:])
		}
		if len(result) > 0 {
			return result
		}
	}
	return defaultValue
}

func parseDuration(s string) time.Duration {
	duration, err := time.ParseDuration(s)
	if err != nil {
		// Default to 15 minutes if parsing fails
		return 15 * time.Minute
	}
	return duration
}

