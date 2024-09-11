package repositories

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/arizdn234/EvoPay/internal/models"
	"github.com/arizdn234/EvoPay/internal/redis"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) Create(user *models.User) error {
	return ur.DB.Create(user).Error
}

func (ur *UserRepository) Update(user *models.User) error {
	return ur.DB.Save(user).Error
}

func (ur *UserRepository) Delete(id uint) error {
	return ur.DB.Delete(&models.User{}, id).Error
}

func (ur *UserRepository) FindAll(user *[]models.User) error {
	return ur.DB.Find(user).Error
}

func (ur *UserRepository) FindByID(id uint) (*models.User, error) {
	cacheKey := "user:" + strconv.FormatUint(uint64(id), 10)
	cachedUser, err := redis.Get(cacheKey)
	if err == nil && cachedUser != "" {
		var user models.User
		if err := json.Unmarshal([]byte(cachedUser), &user); err == nil {
			return &user, nil
		}
	}

	var user models.User
	if err := ur.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	// Cache the user data
	userJSON, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	if err := redis.Set(cacheKey, string(userJSON), time.Minute*10); err != nil {
		log.Printf("Failed to cache user data: %v", err)
	}

	return &user, nil
}

func (ur *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := ur.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) FindRoleByName(name string) (*models.Role, error) {
	var role models.Role
	if err := ur.DB.Where("name = ?", name).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (ur *UserRepository) FindUserByVerificationToken(token string) (*models.User, error) {
	var user models.User
	if err := ur.DB.Where("verification_token = ?", token).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) FindUserByResetToken(token string) (*models.User, error) {
	var user models.User
	if err := ur.DB.Where("reset_token = ?", token).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
