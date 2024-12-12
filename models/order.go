package models

import "time"

type FoodOrder struct {
	ID       uint `json:"id" gorm:"primary_key"`
	OrderID  uint `json:"order_id" gorm:"not null"` // Foreign key to Order
	FoodID   uint `json:"food_id" gorm:"not null"`
	Food     Food `json:"food" gorm:"foreignKey:FoodID"`
	Quantity int  `json:"quantity"`
}

type TableOrder struct {
	ID      uint  `json:"id" gorm:"primary_key"`
	OrderID uint  `json:"order_id" gorm:"not null"` // Foreign key to Order
	TableID uint  `json:"table_id" gorm:"not null"` // Foreign key to Table
	Table   Table `json:"table" gorm:"foreignKey:TableID"`
}

type AddonOrder struct {
	ID       uint  `json:"id" gorm:"primary_key"`
	OrderID  uint  `json:"order_id" gorm:"not null"` // Foreign key to Order
	AddonID  uint  `json:"addon_id" gorm:"not null"` // Foreign key field
	Addon    Addon `json:"addon"`
	Quantity int   `json:"quantity"`
}

type Order struct {
	ID            uint         `json:"id" gorm:"primary_key"`
	UserID        uint         `json:"user_id"`
	RestaurantID  uint         `json:"restaurant_id"`
	TableID       *uint        `json:"table_id,omitempty"`
	FoodOrders    []FoodOrder  `json:"food_orders" gorm:"foreignKey:OrderID"`
	TableOrders   []TableOrder `json:"table_orders" gorm:"foreignKey:OrderID"`
	AddonOrders   []AddonOrder `json:"addon_orders" gorm:"foreignKey:OrderID"`
	TotalPrice    float64      `json:"total_price"`
	Status        string       `json:"status" default:"pending"`
	SpecialNotes  string       `json:"special_notes,omitempty"`
	ExpectedReady *time.Time   `json:"expected_ready,omitempty"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
}
