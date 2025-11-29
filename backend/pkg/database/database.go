package database

import (
	"fmt"

	"github.com/madr/backend/internal/config"
	appLogger "github.com/madr/backend/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

// Connect establishes a connection to the database
func Connect() error {
	dsn := config.AppConfig.Database.GetDSN()

	// Configure GORM logger
	gormLogLevel := gormLogger.Default
	if config.AppConfig.Logging.Level == "debug" {
		gormLogLevel = gormLogger.Default.LogMode(gormLogger.Info)
	} else {
		gormLogLevel = gormLogger.Default.LogMode(gormLogger.Silent)
	}

	// Connect to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogLevel,
	})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	DB = db

	appLogger.Info().
		Str("host", config.AppConfig.Database.Host).
		Str("port", config.AppConfig.Database.Port).
		Str("database", config.AppConfig.Database.Name).
		Msg("Database connected successfully")

	return nil
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}

// HealthCheck performs a simple database health check
func HealthCheck() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("database ping failed: %w", err)
	}

	return nil
}

// Close closes the database connection
func Close() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

