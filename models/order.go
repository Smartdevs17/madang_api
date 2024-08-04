package models

import "time"

type Order struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	UserID       uint      `json:"user_id"`
	RestaurantID uint      `json:"restaurant_id"`
	TableID      *uint     `json:"table_id,omitempty"` // Nullable, as the order might not include a table
	Foods        []Food    `json:"foods" gorm:"many2many:order_foods;"`
	Addons       []Addon   `json:"addons" gorm:"many2many:order_addons;"`
	TotalPrice   float64   `json:"total_price"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
