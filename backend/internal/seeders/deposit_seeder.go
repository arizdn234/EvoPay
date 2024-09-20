package seeders

import (
	"log"

	"github.com/arizdn234/EvoPay/internal/models"
	"gorm.io/gorm"
)

type DepositSeeder struct {
	DB *gorm.DB
}

func (ds *DepositSeeder) Seed() error {
	if err := ds.seedDeposits(); err != nil {
		return err
	}

	return nil
}

func (ds *DepositSeeder) seedDeposits() error {
	var count int64
	if err := ds.DB.Model(&models.Deposit{}).Count(&count).Error; err != nil {
		return err
	}

	// If there is already data in the table, skip seeding
	if count > 0 {
		return nil
	}

	// Get User ID
	var users []models.User
	if err := ds.DB.Find(&users).Error; err != nil {
		return err
	}

	deposits := []models.Deposit{
		{UserID: users[0].ID, Amount: 500.00, DepositMethod: "Credit Card", Status: "completed"},
		{UserID: users[1].ID, Amount: 750.00, DepositMethod: "Bank Transfer", Status: "pending"},
		{UserID: users[2].ID, Amount: 200.00, DepositMethod: "PayPal", Status: "completed"},
		{UserID: users[3].ID, Amount: 100.00, DepositMethod: "Credit Card", Status: "failed"},
	}

	for _, deposit := range deposits {
		if err := ds.DB.Create(&deposit).Error; err != nil {
			return err
		}
		log.Println("Seeded deposit:", deposit)
	}

	return nil
}
