package models

import "time"

type TransactionLog struct {
	ID            uint        `json:"id" gorm:"primaryKey"`
	TransactionID uint        `json:"transaction_id" gorm:"not null"`
	UserID        uint        `json:"user_id" gorm:"not null"`
	Amount        float64     `json:"amount" gorm:"not null"`
	Currency      string      `json:"currency" gorm:"type:varchar(3);not null;default:'USD'"`
	Type          string      `json:"type" gorm:"size:20;not null"`
	Timestamp     time.Time   `json:"timestamp" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	Transaction   Transaction `json:"transaction" gorm:"foreignKey:TransactionID"`
	User          User        `json:"user" gorm:"foreignKey:UserID"`
}
