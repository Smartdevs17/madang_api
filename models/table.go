package models

import "time"

type Table struct {
	ID            uint      `json:"id" gorm:"primary_key"`
	Name          string    `json:"name"`
	Number        int       `json:"number"`
	Capacity      int       `json:"capacity"`
	Image         string    `json:"image"`
	Price         float64   `json:"price"`
	AverageRating float64   `json:"average_rating"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	RestaurantID  uint      `json:"restaurant_id"`
	CategoryId    uint      `json:"category_id"`
	Addons        []Addon   `json:"addons" gorm:"many2many:table_addons;"`
}
