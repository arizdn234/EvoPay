package repositories

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/arizdn234/EvoPay/internal/models"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type DepositRepository interface {
	Create(deposit *models.Deposit) error
	FindByID(id uint) (*models.Deposit, error)
	FindByUserID(userID uint) ([]models.Deposit, error)
	Update(deposit *models.Deposit) error
	Deposit(userID uint, amount float64, currency string) error
}

type depositRepository struct {
	DB                *gorm.DB
	BalanceRepository BalanceRepository
	LogRepository     TransactionLogRepository
	Cache             *redis.Client
}

func NewDepositRepository(db *gorm.DB, balanceRepo BalanceRepository, logRepo TransactionLogRepository, cache *redis.Client) DepositRepository {
	return &depositRepository{
		DB:                db,
		BalanceRepository: balanceRepo,
		LogRepository:     logRepo,
		Cache:             cache,
	}
}

func (r *depositRepository) Create(deposit *models.Deposit) error {
	return r.DB.Create(deposit).Error
}

func (r *depositRepository) FindByID(id uint) (*models.Deposit, error) {
	var deposit models.Deposit
	err := r.DB.Where("id = ?", id).First(&deposit).Error
	return &deposit, err
}

func (r *depositRepository) FindByUserID(userID uint) ([]models.Deposit, error) {
	var deposits []models.Deposit
	err := r.DB.Where("user_id = ?", userID).Find(&deposits).Error
	return deposits, err
}

func (r *depositRepository) Update(deposit *models.Deposit) error {
	return r.DB.Save(deposit).Error
}

// Deposit automation: update balance and insert into TransactionLog
func (r *depositRepository) Deposit(userID uint, amount float64, currency string) error {
	// Validate the deposit amount
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	// Update balance in Redis
	ctx := context.Background()
	if err := r.BalanceRepository.AddToBalance(userID, amount); err != nil {
		return err
	}

	// Cache the updated balance
	userIDStr := strconv.FormatUint(uint64(userID), 10)
	r.Cache.Set(ctx, "balance:"+userIDStr, amount, 0)

	// Create deposit entry
	deposit := &models.Deposit{
		UserID:    userID,
		Amount:    amount,
		Currency:  currency,
		Status:    "completed",
		CreatedAt: time.Now(),
	}
	if err := r.Create(deposit); err != nil {
		return err
	}

	// Create transaction log entry
	return r.createTransactionLog(userID, amount, currency, deposit.ID)
}

// Helper function to create transaction log
func (r *depositRepository) createTransactionLog(userID uint, amount float64, currency string, transactionID uint) error {
	log := &models.TransactionLog{
		UserID:        userID,
		Amount:        amount,
		Currency:      currency,
		TransactionID: transactionID,
		Type:          "deposit",
		Timestamp:     time.Now(),
	}
	return r.LogRepository.Create(log)
}
