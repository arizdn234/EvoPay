package models

import "time"

type Withdrawal struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	UserID           uint      `json:"user_id" gorm:"not null"`
	Amount           float64   `json:"amount" gorm:"not null"`
	Currency         string    `json:"currency" gorm:"type:varchar(3);not null;default:'USD'"`
	WithdrawalMethod string    `json:"withdrawal_method" gorm:"size:50;not null"`
	Status           string    `json:"status" gorm:"size:20;not null"`
	RequestedAt      time.Time `json:"requested_at" gorm:"autoCreateTime"`
	User             User      `json:"user" gorm:"foreignKey:UserID"`
}