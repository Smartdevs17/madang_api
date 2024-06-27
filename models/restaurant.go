package models

type Restaurant struct {
	ID            uint     `json:"id" gorm:"primary_key"`
	Name          string   `json:"name"`
	Address       string   `json:"address"`
	UserID        uint     `json:"user_id"` // Manager's UserID
	Foods         []Food   `json:"foods" gorm:"foreignKey:RestaurantID"`
	Tables        []Table  `json:"tables" gorm:"foreignKey:RestaurantID"`
	Addons        []Addon  `json:"addons" gorm:"foreignKey:RestaurantID"`
	Ratings       []Rating `json:"ratings" gorm:"foreignKey:RestaurantID"`
	AverageRating float64  `json:"average_rating"`
}
