package models

import (
	"time"
)

type Transaction struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	UserID        uint      `json:"user_id" gorm:"not null"`
	User          User      `json:"user" gorm:"foreignKey:UserID"`
	Amount        float64   `json:"amount" gorm:"not null"`
	Currency      string    `json:"currency" gorm:"size:10;not null"`
	PaymentMethod string    `json:"payment_method" gorm:"size:50;not null"`
	Status        string    `json:"status" gorm:"size:20;not null"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
