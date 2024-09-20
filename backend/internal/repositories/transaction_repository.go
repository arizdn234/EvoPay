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

type TransactionRepository interface {
	Create(transaction *models.Transaction) error
	FindAll() ([]models.Transaction, error)
	FindByID(id uint) (*models.Transaction, error)
	FindByUserID(userID uint) ([]models.Transaction, error)
	Update(transaction *models.Transaction) error
	LogTransaction(userID uint, amount float64, currency, paymentMethod, status string) error
}

type transactionRepository struct {
	DB            *gorm.DB
	LogRepository TransactionLogRepository
	Cache         *redis.Client
}

func NewTransactionRepository(db *gorm.DB, logRepo TransactionLogRepository, cache *redis.Client) TransactionRepository {
	return &transactionRepository{
		DB:            db,
		LogRepository: logRepo,
		Cache:         cache,
	}
}

func (r *transactionRepository) Create(transaction *models.Transaction) error {
	return r.DB.Create(transaction).Error
}

func (r *transactionRepository) FindAll() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.DB.Find(&transactions).Error
	return transactions, err
}

func (r *transactionRepository) FindByID(id uint) (*models.Transaction, error) {
	var transaction models.Transaction
	err := r.DB.Where("id = ?", id).First(&transaction).Error
	return &transaction, err
}

func (r *transactionRepository) FindByUserID(userID uint) ([]models.Transaction, error) {
	ctx := context.Background()
	cacheKey := "transactions:user:" + strconv.Itoa(int(userID))

	// Check the cache first
	cachedTransactions, err := r.Cache.Get(ctx, cacheKey).Result()
	if err == nil {
		var transactions []models.Transaction
		if err := json.Unmarshal([]byte(cachedTransactions), &transactions); err == nil {
			return transactions, nil
		}
	}

	// If not found in cache, fetch from DB
	var transactions []models.Transaction
	err = r.DB.Where("user_id = ?", userID).Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	// Cache the result for future requests
	newCachedTransactions, err := json.Marshal(transactions)
	if err != nil {
		return nil, fmt.Errorf("error marshalling transactions: %w", err)
	}

	// Store the marshalled JSON as a string in the cache
	err = r.Cache.Set(ctx, cacheKey, string(newCachedTransactions), 10*time.Minute).Err()
	if err != nil {
		return nil, fmt.Errorf("error setting cache: %w", err)
	}

	return transactions, nil
}

func (r *transactionRepository) Update(transaction *models.Transaction) error {
	return r.DB.Save(transaction).Error
}

// Log transaction to the TransactionLog
func (r *transactionRepository) LogTransaction(userID uint, amount float64, currency, paymentMethod, status string) error {
	transaction := &models.Transaction{
		UserID:        userID,
		Amount:        amount,
		Currency:      currency,
		PaymentMethod: paymentMethod,
		Status:        status,
		CreatedAt:     time.Now(),
	}

	if err := r.Create(transaction); err != nil {
		return err
	}

	// Add to transaction log
	log := &models.TransactionLog{
		UserID:        userID,
		Amount:        amount,
		Currency:      currency,
		TransactionID: transaction.ID,
		Type:          "transaction",
		Timestamp:     time.Now(),
	}
	return r.LogRepository.Create(log)
}
