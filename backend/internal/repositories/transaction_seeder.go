package repositories

import (
	"log"

	"github.com/arizdn234/EvoPay/internal/models"
	"gorm.io/gorm"
)

type TransactionSeeder struct {
	DB *gorm.DB
}

func (ts *TransactionSeeder) Seed() error {
	if err := ts.seedTransactions(); err != nil {
		return err
	}

	return nil
}

func (ts *TransactionSeeder) seedTransactions() error {
	var count int64
	if err := ts.DB.Model(&models.Transaction{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	// Retrieve all users to get their IDs
	var users []models.User
	if err := ts.DB.Find(&users).Error; err != nil {
		return err
	}

	transactions := []models.Transaction{
		{UserID: users[0].ID, Amount: 100.50, Currency: "USD", PaymentMethod: "Credit Card", Status: "completed"},
		{UserID: users[1].ID, Amount: 250.75, Currency: "EUR", PaymentMethod: "PayPal", Status: "pending"},
		{UserID: users[2].ID, Amount: 320.00, Currency: "USD", PaymentMethod: "Bank Transfer", Status: "failed"},
		{UserID: users[3].ID, Amount: 50.00, Currency: "USD", PaymentMethod: "Credit Card", Status: "completed"},
	}

	for _, transaction := range transactions {
		if err := ts.DB.Create(&transaction).Error; err != nil {
			return err
		}
		log.Println("Seeded transaction:", transaction)
	}

	return nil
}
