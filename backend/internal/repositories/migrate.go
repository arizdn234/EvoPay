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

	// Migrate Transaction table
	if err := db.AutoMigrate(&models.Transaction{}); err != nil {
		return err
	}

	// Migrate Balance table
	if err := db.AutoMigrate(&models.Balance{}); err != nil {
		return err
	}

	// Migrate Deposit table
	if err := db.AutoMigrate(&models.Deposit{}); err != nil {
		return err
	}

	// Migrate Withdrawal table
	if err := db.AutoMigrate(&models.Withdrawal{}); err != nil {
		return err
	}

	// Migrate TransactionLog table
	if err := db.AutoMigrate(&models.TransactionLog{}); err != nil {
		return err
	}

	return nil
}
