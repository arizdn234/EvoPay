package repositories

import (
	"context"
	"strconv"

	"github.com/arizdn234/EvoPay/internal/models"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type TransactionLogRepository interface {
	Create(log *models.TransactionLog) error
	FindByTransactionID(transactionID uint) ([]models.TransactionLog, error)
	FindByUserID(userID uint) ([]models.TransactionLog, error)
}

type transactionLogRepository struct {
	DB    *gorm.DB
	Cache *redis.Client
}

func NewTransactionLogRepository(db *gorm.DB, cache *redis.Client) TransactionLogRepository {
	return &transactionLogRepository{DB: db, Cache: cache}
}

func (r *transactionLogRepository) Create(log *models.TransactionLog) error {
	return r.DB.Create(log).Error
}

func (r *transactionLogRepository) FindByTransactionID(transactionID uint) ([]models.TransactionLog, error) {
	ctx := context.Background()
	cacheKey := "transaction_logs:transaction:" + strconv.FormatUint(uint64(transactionID), 10)

	var logs []models.TransactionLog
	if err := r.Cache.Get(ctx, cacheKey).Scan(&logs); err == nil {
		return logs, nil // Return cached logs if found
	}

	// If not found in cache, retrieve from DB
	err := r.DB.Where("transaction_id = ?", transactionID).Find(&logs).Error
	if err != nil {
		return logs, err
	}

	// Cache the result for future queries
	r.Cache.Set(ctx, cacheKey, logs, 0) // Use an appropriate TTL if necessary
	return logs, nil
}

func (r *transactionLogRepository) FindByUserID(userID uint) ([]models.TransactionLog, error) {
	ctx := context.Background()
	cacheKey := "transaction_logs:user:" + strconv.FormatUint(uint64(userID), 10)

	var logs []models.TransactionLog
	if err := r.Cache.Get(ctx, cacheKey).Scan(&logs); err == nil {
		return logs, nil // Return cached logs if found
	}

	// If not found in cache, retrieve from DB
	err := r.DB.Where("user_id = ?", userID).Find(&logs).Error
	if err != nil {
		return logs, err
	}

	// Cache the result for future queries
	r.Cache.Set(ctx, cacheKey, logs, 0) // Use an appropriate TTL if necessary
	return logs, nil
}
