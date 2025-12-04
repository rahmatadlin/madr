package seed

import (
	donationCategoryDomain "github.com/madr/backend/internal/domain/donationcategory"
	"github.com/madr/backend/internal/domain/user"
	donationCategoryRepo "github.com/madr/backend/internal/repository/donationcategory"
	userRepo "github.com/madr/backend/internal/repository/user"
	"github.com/madr/backend/pkg/bcrypt"
	"github.com/madr/backend/pkg/logger"
)

// SeedDefaultAdmin seeds a default admin user
func SeedDefaultAdmin() error {
	repo := userRepo.NewRepository()

	// Check if admin already exists
	exists, err := repo.ExistsByUsername("admin")
	if err != nil {
		return err
	}
	if exists {
		logger.Info().Msg("Default admin already exists, skipping seed")
		return nil
	}

	// Create default admin
	hashedPassword, err := bcrypt.HashPassword("admin123")
	if err != nil {
		return err
	}

	adminUser := &user.User{
		Username: "admin",
		Email:    "admin@madr.local",
		Password: hashedPassword,
		Name:     "Default Admin",
		Role:     user.RoleAdmin,
		IsActive: true,
	}

	if err := repo.Create(adminUser); err != nil {
		return err
	}

	logger.Info().
		Str("username", "admin").
		Str("email", "admin@madr.local").
		Msg("Default admin user created successfully")

	return nil
}

// SeedDonationCategories seeds default donation categories
func SeedDonationCategories() error {
	repo := donationCategoryRepo.NewRepository()

	categories := []donationCategoryDomain.DonationCategory{
		{
			Name:        "Pembangunan",
			Description: "Donasi untuk pembangunan masjid",
		},
		{
			Name:        "Operasional",
			Description: "Donasi untuk operasional masjid",
		},
		{
			Name:        "Sosial",
			Description: "Donasi untuk kegiatan sosial",
		},
		{
			Name:        "Anak Yatim",
			Description: "Donasi untuk program anak yatim",
		},
	}

	for _, cat := range categories {
		// Check if category already exists
		exists, err := repo.ExistsByName(cat.Name, 0)
		if err != nil {
			logger.Warn().Err(err).Str("category", cat.Name).Msg("Failed to check category existence")
			continue
		}
		if exists {
			logger.Info().Str("category", cat.Name).Msg("Category already exists, skipping")
			continue
		}

		// Create category
		if err := repo.Create(&cat); err != nil {
			logger.Warn().Err(err).Str("category", cat.Name).Msg("Failed to seed category")
			continue
		}

		logger.Info().Str("category", cat.Name).Msg("Donation category seeded successfully")
	}

	return nil
}

// SeedAll runs all seeders
func SeedAll() error {
	logger.Info().Msg("Starting database seeding...")

	if err := SeedDefaultAdmin(); err != nil {
		logger.Warn().Err(err).Msg("Failed to seed default admin")
		// Don't return error, continue with other seeds
	}

	if err := SeedDonationCategories(); err != nil {
		logger.Warn().Err(err).Msg("Failed to seed donation categories")
		// Don't return error, continue
	}

	logger.Info().Msg("Database seeding completed")
	return nil
}

