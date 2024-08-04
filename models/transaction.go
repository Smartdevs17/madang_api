package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	OrderID      uint      `json:"order_id"`
	PaymentID    uint      `json:"payment_id"`
	Status       string    `json:"status"` // e.g., "initiated", "completed", "failed"
	Amount       float64   `json:"amount"`
	RestaurantID uint      `json:"restaurant_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
