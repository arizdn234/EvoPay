package models

import "time"

type Role struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(255);unique;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID                uint       `json:"id" gorm:"primaryKey"`
	Name              string     `json:"name" gorm:"type:varchar(255);"`
	Email             string     `json:"email" gorm:"type:varchar(255);not null"`
	Password          string     `json:"password" gorm:"type:varchar(255);not null"`
	RoleID            uint       `json:"role_id" gorm:"not null"`
	Role              Role       `json:"role" gorm:"foreignKey:RoleID"`
	ResetToken        *string    `json:"-" gorm:"unique"`
	ResetTokenExpiry  *time.Time `json:"-"`
	EmailVerified     bool       `json:"email_verified"`
	VerificationToken *string    `json:"-"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRegister struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
