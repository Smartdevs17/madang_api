package models

type Table struct {
	ID           uint    `json:"id" gorm:"primary_key"`
	Name         string  `json:"name"`
	Number       int     `json:"number"`
	Capacity     int     `json:"capacity"`
	Image        string  `json:"image"`
	Price        float64 `json:"price"`
	RestaurantID uint    `json:"restaurant_id"`
	CategoryId   uint    `json:"category_id"`
	Addons       []Addon `json:"addons" gorm:"many2many:table_addons;"`
}
