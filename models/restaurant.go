package models

type Restaurant struct {
	ID            uint     `json:"id" gorm:"primary_key"`
	Name          string   `json:"name"`
	Address       string   `json:"address"`
	UserID        uint     `json:"user_id"` // Manager's UserID
	Phone         string   `json:"phone"`
	Email         string   `json:"email"`
	Website       string   `json:"website"`
	Location      string   `json:"location"`
	State         string   `json:"state"`
	Country       string   `json:"country"`
	Image         string   `json:"image"`
	OpeningHours  string   `json:"opening_hours"`
	ClosingHours  string   `json:"closing_hours"`
	Active        bool     `json:"active"`
	Verified      bool     `json:"verified"`
	VerfiedAt     string   `json:"verified_at"`
	Foods         []Food   `json:"foods" gorm:"foreignKey:RestaurantID"`
	Tables        []Table  `json:"tables" gorm:"foreignKey:RestaurantID"`
	Addons        []Addon  `json:"addons" gorm:"foreignKey:RestaurantID"`
	Ratings       []Rating `json:"ratings" gorm:"foreignKey:RestaurantID"`
	AverageRating float64  `json:"average_rating"`
}
