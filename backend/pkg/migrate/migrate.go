package migrate

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/madr/backend/internal/config"
	"github.com/madr/backend/pkg/logger"
)

// getMigrationsPath returns the absolute path to migrations directory
func getMigrationsPath() (string, error) {
	// Try to get migrations path relative to current working directory
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Check if migrations folder exists in current directory
	migrationsPath := filepath.Join(wd, "migrations")
	if _, err := os.Stat(migrationsPath); err == nil {
		return migrationsPath, nil
	}

	// Try backend/migrations (if running from root)
	migrationsPath = filepath.Join(wd, "backend", "migrations")
	if _, err := os.Stat(migrationsPath); err == nil {
		return migrationsPath, nil
	}

	return "", fmt.Errorf("migrations directory not found")
}

// RunMigrations runs all pending migrations
func RunMigrations(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create postgres driver: %w", err)
	}

	// Get migrations directory path
	migrationsPath, err := getMigrationsPath()
	if err != nil {
		return fmt.Errorf("failed to find migrations directory: %w", err)
	}

	migrationsURL := fmt.Sprintf("file://%s", migrationsPath)

	m, err := migrate.NewWithDatabaseInstance(
		migrationsURL,
		config.AppConfig.Database.Name,
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	// Run migrations
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			logger.Info().Msg("No pending migrations")
			return nil
		}
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	logger.Info().Msg("Database migrations completed successfully")
	return nil
}

// DownMigrations rolls back the last migration
func DownMigrations(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create postgres driver: %w", err)
	}

	migrationsPath, err := getMigrationsPath()
	if err != nil {
		return fmt.Errorf("failed to find migrations directory: %w", err)
	}

	migrationsURL := fmt.Sprintf("file://%s", migrationsPath)

	m, err := migrate.NewWithDatabaseInstance(
		migrationsURL,
		config.AppConfig.Database.Name,
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	if err := m.Down(); err != nil {
		if err == migrate.ErrNoChange {
			logger.Info().Msg("No migrations to rollback")
			return nil
		}
		return fmt.Errorf("failed to rollback migration: %w", err)
	}

	logger.Info().Msg("Migration rolled back successfully")
	return nil
}

// GetMigrationVersion returns the current migration version
func GetMigrationVersion(db *sql.DB) (uint, bool, error) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return 0, false, fmt.Errorf("failed to create postgres driver: %w", err)
	}

	migrationsPath, err := getMigrationsPath()
	if err != nil {
		return 0, false, fmt.Errorf("failed to find migrations directory: %w", err)
	}

	migrationsURL := fmt.Sprintf("file://%s", migrationsPath)

	m, err := migrate.NewWithDatabaseInstance(
		migrationsURL,
		config.AppConfig.Database.Name,
		driver,
	)
	if err != nil {
		return 0, false, fmt.Errorf("failed to create migrate instance: %w", err)
	}

	version, dirty, err := m.Version()
	if err != nil {
		if err == migrate.ErrNilVersion {
			return 0, false, nil
		}
		return 0, false, err
	}

	return version, dirty, nil
}

