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
		command = flag.String("command", "up", "Migration command: up, down, version, force")
		version = flag.Int("version", 0, "Version number (required for force command)")
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
		currentVersion, dirty, err := migrate.GetMigrationVersion(sqlDB)
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to get migration version")
			os.Exit(1)
		}
		if dirty {
			logger.Warn().Uint("version", currentVersion).Msg("Database is in dirty state. Use 'force' command to fix.")
		} else {
			logger.Info().Uint("version", currentVersion).Msg("Current migration version")
		}
	case "force":
		if *version < 0 {
			fmt.Fprintf(os.Stderr, "Version number is required for force command\n")
			fmt.Fprintf(os.Stderr, "Usage: migrate -command=force -version=<version_number>\n")
			fmt.Fprintf(os.Stderr, "Example: migrate -command=force -version=5\n")
			os.Exit(1)
		}
		if err := migrate.ForceVersion(sqlDB, *version); err != nil {
			logger.Fatal().Err(err).Msg("Failed to force version")
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", *command)
		fmt.Fprintf(os.Stderr, "Usage: migrate -command=[up|down|version|force]\n")
		fmt.Fprintf(os.Stderr, "For force: migrate -command=force -version=<version_number>\n")
		os.Exit(1)
	}
}

