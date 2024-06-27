package models

type Addon struct {
	ID           uint    `json:"id" gorm:"primary_key"`
	Name         string  `json:"name"`
	Type         string  `json:"type"` // e.g., "chair", "flower"
	Price        float64 `json:"price"`
	RestaurantID uint    `json:"restaurant_id"`
}
