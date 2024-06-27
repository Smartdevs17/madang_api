package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	OrderID uint    `json:"order_id"`
	Amount  float64 `json:"amount"`
	Method  string  `json:"method"` // e.g., "credit_card", "paypal"
	Status  string  `json:"status"` // e.g., "pending", "completed", "failed"
}
