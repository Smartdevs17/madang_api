package models

import "gorm.io/gorm"

type Rating struct {
	gorm.Model
	UserID       uint       `json:"user_id"`
	FoodID       *uint      `json:"food_id,omitempty"`       // Nullable, as the rating might not be for a food item
	RestaurantID *uint      `json:"restaurant_id,omitempty"` // Nullable, as the rating might not be for a restaurant
	Score        int        `json:"score"`                   // Rating score (e.g., 1-5)
	Comment      string     `json:"comment"`
	User         User       `json:"user"`
	Food         Food       `json:"food"`
	Restaurant   Restaurant `json:"restaurant"`
}
