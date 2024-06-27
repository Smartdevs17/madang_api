package models

type Table struct {
	ID           uint    `json:"id" gorm:"primary_key"`
	Number       int     `json:"number"`
	Capacity     int     `json:"capacity"`
	RestaurantID uint    `json:"restaurant_id"`
	Addons       []Addon `json:"addons" gorm:"many2many:table_addons;"`
}
