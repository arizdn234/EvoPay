package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/arizdn234/EvoPay/internal/models"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type WithdrawalRepository interface {
	Create(withdrawal *models.Withdrawal) error
	FindByID(id uint) (*models.Withdrawal, error)
	FindByUserID(userID uint) ([]models.Withdrawal, error)
	Update(withdrawal *models.Withdrawal) error
	Withdraw(userID uint, amount float64, currency string) error
}

type withdrawalRepository struct {
	DB                *gorm.DB
	BalanceRepository BalanceRepository
	LogRepository     TransactionLogRepository
	Cache             *redis.Client
}

func NewWithdrawalRepository(db *gorm.DB, balanceRepo BalanceRepository, logRepo TransactionLogRepository, cache *redis.Client) WithdrawalRepository {
	return &withdrawalRepository{
		DB:                db,
		BalanceRepository: balanceRepo,
		LogRepository:     logRepo,
		Cache:             cache,
	}
}

func (r *withdrawalRepository) Create(withdrawal *models.Withdrawal) error {
	return r.DB.Create(withdrawal).Error
}

func (r *withdrawalRepository) FindByID(id uint) (*models.Withdrawal, error) {
	var withdrawal models.Withdrawal
	err := r.DB.Where("id = ?", id).First(&withdrawal).Error
	return &withdrawal, err
}

func (r *withdrawalRepository) FindByUserID(userID uint) ([]models.Withdrawal, error) {
	ctx := context.Background()
	cacheKey := "withdrawals:user:" + strconv.Itoa(int(userID))

	// Check the cache first
	cachedWithdrawals, err := r.Cache.Get(ctx, cacheKey).Result()
	if err == nil {
		var withdrawals []models.Withdrawal
		if err := json.Unmarshal([]byte(cachedWithdrawals), &withdrawals); err == nil {
			return withdrawals, nil
		}
	}

	// If not found in cache, fetch from DB
	var withdrawals []models.Withdrawal
	err = r.DB.Where("user_id = ?", userID).Find(&withdrawals).Error
	if err != nil {
		return nil, err
	}

	// Cache the result for future requests
	newCachedWithdrawals, err := json.Marshal(withdrawals)
	if err != nil {
		return nil, fmt.Errorf("error marshalling withdrawals: %w", err)
	}

	// Store the marshalled JSON as a string in the cache
	err = r.Cache.Set(ctx, cacheKey, string(newCachedWithdrawals), 10*time.Minute).Err()
	if err != nil {
		return nil, fmt.Errorf("error setting cache: %w", err)
	}

	return withdrawals, nil
}

func (r *withdrawalRepository) Update(withdrawal *models.Withdrawal) error {
	return r.DB.Save(withdrawal).Error
}

// Withdrawal automation: update balance and insert into TransactionLog
func (r *withdrawalRepository) Withdraw(userID uint, amount float64, currency string) error {
	// Reduce user balance
	if err := r.BalanceRepository.SubtractFromBalance(userID, amount); err != nil {
		return err
	}

	// Create withdrawal entry
	withdrawal := &models.Withdrawal{
		UserID:      userID,
		Amount:      amount,
		Currency:    currency,
		Status:      "completed",
		RequestedAt: time.Now(),
	}
	if err := r.Create(withdrawal); err != nil {
		return err
	}

	// Add to transaction log
	log := &models.TransactionLog{
		UserID:        userID,
		Amount:        amount,
		Currency:      currency,
		TransactionID: withdrawal.ID,
		Type:          "withdrawal",
		Timestamp:     time.Now(),
	}
	return r.LogRepository.Create(log)
}
