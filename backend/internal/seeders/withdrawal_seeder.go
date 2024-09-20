package seeders

import (
	"log"

	"github.com/arizdn234/EvoPay/internal/models"
	"gorm.io/gorm"
)

type WithdrawalSeeder struct {
	DB *gorm.DB
}

func (ws *WithdrawalSeeder) Seed() error {
	if err := ws.seedWithdrawals(); err != nil {
		return err
	}

	return nil
}

func (ws *WithdrawalSeeder) seedWithdrawals() error {
	var count int64
	if err := ws.DB.Model(&models.Withdrawal{}).Count(&count).Error; err != nil {
		return err
	}

	// If there is already data in the table, skip seeding
	if count > 0 {
		return nil
	}

	// Get User ID
	var users []models.User
	if err := ws.DB.Find(&users).Error; err != nil {
		return err
	}

	withdrawals := []models.Withdrawal{
		{UserID: users[0].ID, Amount: 300.00, WithdrawalMethod: "Bank Transfer", Status: "completed"},
		{UserID: users[1].ID, Amount: 150.00, WithdrawalMethod: "PayPal", Status: "pending"},
		{UserID: users[2].ID, Amount: 400.00, WithdrawalMethod: "Bank Transfer", Status: "failed"},
		{UserID: users[3].ID, Amount: 100.00, WithdrawalMethod: "Credit Card", Status: "completed"},
	}

	for _, withdrawal := range withdrawals {
		if err := ws.DB.Create(&withdrawal).Error; err != nil {
			return err
		}
		log.Println("Seeded withdrawal:", withdrawal)
	}

	return nil
}
