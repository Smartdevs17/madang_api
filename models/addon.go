package models

import "time"

type Addon struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	Name         string    `json:"name"`
	Type         string    `json:"type"` // e.g., "chair", "flower"
	Price        float64   `json:"price"`
	RestaurantID uint      `json:"restaurant_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
