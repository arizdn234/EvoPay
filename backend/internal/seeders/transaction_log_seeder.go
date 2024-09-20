package seeders

import (
	logSys "log"
	"time"

	"github.com/arizdn234/EvoPay/internal/models"
	"gorm.io/gorm"
)

type TransactionLogSeeder struct {
	DB *gorm.DB
}

func (tls *TransactionLogSeeder) Seed() error {
	if err := tls.seedTransactionLogs(); err != nil {
		return err
	}

	return nil
}

func (tls *TransactionLogSeeder) seedTransactionLogs() error {
	var count int64
	if err := tls.DB.Model(&models.TransactionLog{}).Count(&count).Error; err != nil {
		return err
	}

	// If there is already data in the table, skip seeding
	if count > 0 {
		return nil
	}

	// Get Transaction ID
	var transactions []models.Transaction
	if err := tls.DB.Find(&transactions).Error; err != nil {
		return err
	}

	transactionLogs := []models.TransactionLog{
		{
			TransactionID: transactions[0].ID,
			UserID:        transactions[0].UserID,
			Amount:        transactions[0].Amount,
			Currency:      transactions[0].Currency,
			Type:          "transaction",
			Timestamp:     time.Now(),
		},
		{
			TransactionID: transactions[1].ID,
			UserID:        transactions[1].UserID,
			Amount:        transactions[1].Amount,
			Currency:      transactions[1].Currency,
			Type:          "transaction",
			Timestamp:     time.Now(),
		},
		{
			TransactionID: transactions[2].ID,
			UserID:        transactions[2].UserID,
			Amount:        transactions[2].Amount,
			Currency:      transactions[2].Currency,
			Type:          "transaction",
			Timestamp:     time.Now(),
		},
		{
			TransactionID: transactions[3].ID,
			UserID:        transactions[3].UserID,
			Amount:        transactions[3].Amount,
			Currency:      transactions[3].Currency,
			Type:          "transaction",
			Timestamp:     time.Now(),
		},
	}

	for _, log := range transactionLogs {
		if err := tls.DB.Create(&log).Error; err != nil {
			return err
		}
		logSys.Println("Seeded transaction log:", log)
	}

	return nil
}
