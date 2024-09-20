package seeders

import (
	"log"

	"gorm.io/gorm"
)

// Seeder function
func SeedAll(db *gorm.DB) error {
	// User seeder
	userSeeder := &UserSeeder{DB: db}
	if err := userSeeder.Seed(); err != nil {
		log.Println("Failed to seed users: ", err)
		return err
	}

	// Transaction seeder
	transactionSeeder := &TransactionSeeder{DB: db}
	if err := transactionSeeder.Seed(); err != nil {
		log.Println("Failed to seed transactions: ", err)
		return err
	}

	// Balance seeder
	balanceSeeder := &BalanceSeeder{DB: db}
	if err := balanceSeeder.Seed(); err != nil {
		log.Println("Failed to seed balances: ", err)
		return err
	}

	// Deposit seeder
	depositSeeder := &DepositSeeder{DB: db}
	if err := depositSeeder.Seed(); err != nil {
		log.Println("Failed to seed deposits: ", err)
		return err
	}

	// Withdrawal seeder
	withdrawalSeeder := &WithdrawalSeeder{DB: db}
	if err := withdrawalSeeder.Seed(); err != nil {
		log.Println("Failed to seed withdrawals: ", err)
		return err
	}

	// TransactionLog seeder
	transactionLogSeeder := &TransactionLogSeeder{DB: db}
	if err := transactionLogSeeder.Seed(); err != nil {
		log.Println("Failed to seed transaction logs: ", err)
		return err
	}

	log.Println("All seeders ran successfully")
	return nil
}
