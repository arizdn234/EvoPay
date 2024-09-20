package seeders

import (
	"log"
	"time"

	"github.com/arizdn234/EvoPay/internal/models"
	"gorm.io/gorm"
)

type BalanceSeeder struct {
	DB *gorm.DB
}

func (bs *BalanceSeeder) Seed() error {
	if err := bs.seedBalances(); err != nil {
		return err
	}

	return nil
}

func (bs *BalanceSeeder) seedBalances() error {
	var count int64
	if err := bs.DB.Model(&models.Balance{}).Count(&count).Error; err != nil {
		return err
	}

	// If there is already data in the table, skip seeding
	if count > 0 {
		return nil
	}

	// Get User ID
	var users []models.User
	if err := bs.DB.Find(&users).Error; err != nil {
		return err
	}

	balances := []models.Balance{
		{UserID: users[0].ID, CurrentBalance: 1000, Currency: "USD", LastUpdated: time.Now()},
		{UserID: users[1].ID, CurrentBalance: 500, Currency: "EUR", LastUpdated: time.Now()},
		{UserID: users[2].ID, CurrentBalance: 200, Currency: "JPY", LastUpdated: time.Now()},
	}

	for _, balance := range balances {
		if err := bs.DB.Create(&balance).Error; err != nil {
			return err
		}
		log.Println("Seeded balance:", balance)
	}

	return nil
}
