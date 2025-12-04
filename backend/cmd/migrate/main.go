package main

import (
	"flag"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/madr/backend/internal/config"
	"github.com/madr/backend/pkg/database"
	"github.com/madr/backend/pkg/logger"
	"github.com/madr/backend/pkg/migrate"
)

func main() {
	var (
		command = flag.String("command", "up", "Migration command: up, down, version")
	)
	flag.Parse()

	// Load configuration
	if err := config.Load(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load configuration: %v\n", err)
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

	// Get underlying sql.DB
	sqlDB, err := database.GetDB().DB()
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to get database instance")
		os.Exit(1)
	}

	switch *command {
	case "up":
		if err := migrate.RunMigrations(sqlDB); err != nil {
			logger.Fatal().Err(err).Msg("Migration failed")
			os.Exit(1)
		}
	case "down":
		if err := migrate.DownMigrations(sqlDB); err != nil {
			logger.Fatal().Err(err).Msg("Rollback failed")
			os.Exit(1)
		}
	case "version":
		version, dirty, err := migrate.GetMigrationVersion(sqlDB)
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to get migration version")
			os.Exit(1)
		}
		if dirty {
			logger.Warn().Uint("version", version).Msg("Database is in dirty state")
		} else {
			logger.Info().Uint("version", version).Msg("Current migration version")
		}
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", *command)
		fmt.Fprintf(os.Stderr, "Usage: migrate -command=[up|down|version]\n")
		os.Exit(1)
	}
}

