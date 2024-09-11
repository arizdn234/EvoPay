// repositories/user_seeder.go

package repositories

import (
	"log"

	"github.com/arizdn234/EvoPay/internal/models"
	"github.com/arizdn234/EvoPay/utils"
	"gorm.io/gorm"
)

type UserSeeder struct {
	DB *gorm.DB
}

func (us *UserSeeder) Seed() error {
	// Seed roles first
	if err := us.seedRoles(); err != nil {
		return err
	}

	// Seed users
	if err := us.seedUsers(); err != nil {
		return err
	}

	return nil
}

func (us *UserSeeder) seedRoles() error {
	var count int64
	if err := us.DB.Model(&models.Role{}).Count(&count).Error; err != nil {
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
		if err := us.DB.Create(&role).Error; err != nil {
			return err
		}
		log.Println("Seeded role:", role)
	}

	return nil
}

func (us *UserSeeder) seedUsers() error {
	var count int64
	if err := us.DB.Model(&models.User{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	// Retrieve all roles to get their IDs
	var roles []models.Role
	if err := us.DB.Find(&roles).Error; err != nil {
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
		if err := us.DB.Create(user).Error; err != nil {
			return err
		}
		log.Println("Seeded user:", user)
	}

	return nil
}
