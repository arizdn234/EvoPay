package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/arizdn234/EvoPay/internal/models"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type BalanceRepository interface {
	Create(balance *models.Balance) error
	FindByUserID(userID uint) (*models.Balance, error)
	Update(balance *models.Balance) error
	AddToBalance(userID uint, amount float64) error
	SubtractFromBalance(userID uint, amount float64) error
}

type balanceRepository struct {
	DB    *gorm.DB
	Cache *redis.Client
	Ctx   context.Context
}

func NewBalanceRepository(db *gorm.DB, cache *redis.Client) BalanceRepository {
	return &balanceRepository{DB: db, Cache: cache, Ctx: context.Background()}
}

func (r *balanceRepository) Create(balance *models.Balance) error {
	err := r.DB.Create(balance).Error
	if err == nil {
		// Set cache when creating a new balance
		r.Cache.Set(r.Ctx, getBalanceCacheKey(balance.UserID), balance.CurrentBalance, 0)
	}
	return err
}

func (r *balanceRepository) FindByUserID(userID uint) (*models.Balance, error) {
	cachedBalance, err := r.Cache.Get(r.Ctx, getBalanceCacheKey(userID)).Result()
	if err == nil {
		var balance models.Balance
		if err := json.Unmarshal([]byte(cachedBalance), &balance); err == nil {
			return &balance, nil
		} else {
			fmt.Printf("Error unmarshalling cached balance: %v\n", err)
		}
	}

	var balance models.Balance
	err = r.DB.Where("user_id = ?", userID).First(&balance).Error
	if err != nil {
		return nil, err
	}

	balanceJSON, err := json.Marshal(balance)
	if err == nil {
		r.Cache.Set(r.Ctx, getBalanceCacheKey(userID), balanceJSON, 5*time.Minute)
	} else {
		fmt.Printf("Error marshalling balance for cache: %v\n", err)
	}

	return &balance, nil
}

func (r *balanceRepository) Update(balance *models.Balance) error {
	err := r.DB.Save(balance).Error
	if err == nil {
		r.Cache.Set(r.Ctx, getBalanceCacheKey(balance.UserID), balance.CurrentBalance, 0)
	}
	return err
}

func (r *balanceRepository) AddToBalance(userID uint, amount float64) error {
	balance, err := r.FindByUserID(userID)
	if err != nil {
		return err
	}

	balance.CurrentBalance += amount
	return r.Update(balance)
}

func (r *balanceRepository) SubtractFromBalance(userID uint, amount float64) error {
	balance, err := r.FindByUserID(userID)
	if err != nil {
		return err
	}

	if balance.CurrentBalance < amount {
		return errors.New("insufficient funds")
	}

	balance.CurrentBalance -= amount
	return r.Update(balance)
}

// Helper function to get cache key
func getBalanceCacheKey(userID uint) string {
	return "balance:" + strconv.Itoa(int(userID))
}

// Helper function to parse float from string
// func parseFloat(s string) float64 {
// 	val, _ := strconv.ParseFloat(s, 64)
// 	return val
// }
