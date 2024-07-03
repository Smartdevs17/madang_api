package models

type Table struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Number   int    `json:"number"`
	Capacity int    `json:"capacity"`

	RestaurantID uint    `json:"restaurant_id"`
	CategoryId   uint    `json:"category_id"`
	Addons       []Addon `json:"addons" gorm:"many2many:table_addons;"`
}
