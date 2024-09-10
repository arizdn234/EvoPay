package repository

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/arizdn234/EvoPay/internal/models"
	"github.com/arizdn234/EvoPay/internal/redis"
	"github.com/arizdn234/EvoPay/utils"
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

func Migrate(db *gorm.DB) error {
	// Migrate User table
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return err
	}

	// Migrate Role table
	if err := db.AutoMigrate(&models.Role{}); err != nil {
		return err
	}

	// Migrate Permission table
	if err := db.AutoMigrate(&models.Permission{}); err != nil {
		return err
	}

	// Migrate RolePermission table
	if err := db.AutoMigrate(&models.RolePermission{}); err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) Seed() error {
	// Seed roles first
	if err := ur.seedRoles(); err != nil {
		return err
	}

	// Seed users
	if err := ur.seedUsers(); err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) seedRoles() error {
	var count int64
	if err := ur.DB.Model(&models.Role{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	roles := []models.Role{
		{Name: "admin"},
		{Name: "user"},
	}

	for _, role := range roles {
		if err := ur.DB.Create(&role).Error; err != nil {
			return err
		}
		log.Println("Seeded role:", role)
	}

	return nil
}

func (ur *UserRepository) seedUsers() error {
	var count int64
	if err := ur.DB.Model(&models.User{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	// Retrieve all roles to get their IDs
	var roles []models.Role
	if err := ur.DB.Find(&roles).Error; err != nil {
		return err
	}

	// Map role names to IDs
	roleMap := make(map[string]uint)
	for _, role := range roles {
		roleMap[role.Name] = role.ID
	}

	// Define users with associated role IDs
	users := []*models.User{
		{Name: "Admin", Email: "admin@admin.com", Password: "adminpasS1", RoleID: roleMap["admin"]},
		{Name: "Johnny Deep", Email: "jdeep557@gmail.com", Password: "Blabla2421", RoleID: roleMap["user"]},
		{Name: "Tom Cruise", Email: "tcruise557@gmail.com", Password: "baDls31dw", RoleID: roleMap["user"]},
		{Name: "Billie Elish", Email: "bilelish557@gmail.com", Password: "jup752sDdka", RoleID: roleMap["user"]},
	}

	for _, user := range users {
		// Hash the password
		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			return err
		}
		user.Password = hashedPassword

		user.ResetToken = nil
		user.ResetTokenExpiry = nil
		user.VerificationToken = nil

		// Create the user
		if err := ur.DB.Create(user).Error; err != nil {
			return err
		}
		log.Println("Seeded user:", user)
	}

	return nil
}
