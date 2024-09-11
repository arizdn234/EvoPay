package repositories

import (
	"github.com/arizdn234/EvoPay/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	// Migrate User table
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return err
	}

	// Migrate Role table
	if err := db.AutoMigrate(&models.Role{}); err != nil {
		return err
	}

	// Migrate Permission table
	if err := db.AutoMigrate(&models.Permission{}); err != nil {
		return err
	}

	// Migrate RolePermission table
	if err := db.AutoMigrate(&models.RolePermission{}); err != nil {
		return err
	}

	return nil
}
