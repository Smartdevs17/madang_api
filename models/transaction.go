package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	OrderID      uint    `json:"order_id"`
	PaymentID    uint    `json:"payment_id"`
	Status       string  `json:"status"` // e.g., "initiated", "completed", "failed"
	Amount       float64 `json:"amount"`
	RestaurantID uint    `json:"restaurant_id"`
}
