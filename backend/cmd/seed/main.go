package main

import (
	"os"

	"github.com/madr/backend/internal/config"
	"github.com/madr/backend/pkg/database"
	"github.com/madr/backend/pkg/logger"
	"github.com/madr/backend/pkg/seed"
)

func main() {
	// Load configuration
	if err := config.Load(); err != nil {
		logger.Fatal().Err(err).Msg("Failed to load configuration")
		os.Exit(1)
	}

	// Initialize logger
	logger.Init(config.AppConfig.Logging.Level, config.AppConfig.Logging.Format)

	// Connect to database
	if err := database.Connect(); err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to database")
		os.Exit(1)
	}
	defer database.Close()

	// Run all seeders
	if err := seed.SeedAll(); err != nil {
		logger.Fatal().Err(err).Msg("Seeding failed")
		os.Exit(1)
	}

	logger.Info().Msg("Seeding completed successfully")
}

