package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name         string    `gorm:"not null" json:"name"`
	Type         string    `gorm:"not null" json:"type"`
	RestaurantID uint      `json:"restaurant_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
