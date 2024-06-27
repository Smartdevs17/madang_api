package models

type Food struct {
	ID            uint     `json:"id" gorm:"primary_key"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	Price         float64  `json:"price"`
	RestaurantID  uint     `json:"restaurant_id"`
	Ratings       []Rating `json:"ratings" gorm:"foreignKey:FoodID"`
	AverageRating float64  `json:"average_rating"`
}
