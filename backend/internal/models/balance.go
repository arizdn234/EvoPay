package models

import "time"

type Balance struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	UserID         uint      `json:"user_id" gorm:"not null"`
	CurrentBalance float64   `json:"current_balance" gorm:"not null;default:0"`
	Currency       string    `json:"currency" gorm:"type:varchar(3);not null;default:'USD'"`
	LastUpdated    time.Time `json:"last_updated" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	User           User      `json:"user" gorm:"foreignKey:UserID"`
}
