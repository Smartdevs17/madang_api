package models

import "time"

type Food struct {
	ID            uint      `json:"id" gorm:"primary_key"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Image         string    `json:"image"`
	Price         float64   `json:"price"`
	RestaurantID  uint      `json:"restaurant_id"`
	CategoryId    uint      `json:"category_id"`
	Ratings       []Rating  `json:"ratings" gorm:"foreignKey:FoodID"`
	AverageRating float64   `json:"average_rating"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
