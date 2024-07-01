package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name         string `gorm:"unique;not null" json:"name"`
	Type         string `gorm:"not null" json:"type"`
	RestaurantID uint   `json:"restaurant_id"`
}
