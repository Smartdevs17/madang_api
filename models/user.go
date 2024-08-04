package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name                 string       `json:"name"`
	Email                string       `json:"email"`
	Password             string       `json:"password"`
	Phone                string       `json:"phone"`
	Role                 string       `json:"role"` // "customer" or "manager" or "admin"
	Avatar               string       `json:"avatar"`
	Active               bool         `json:"active"`
	Token                string       `json:"token"`
	DeviceId             string       `json:"device_id"`
	DeviceToken          string       `json:"device_token"`
	EmailVerified        bool         `json:"email_verified"`
	EmailVerificationOTP string       `json:"-"`
	Restaurants          []Restaurant `json:"restaurants" gorm:"foreignKey:UserID"`
	Orders               []Order      `json:"orders" gorm:"foreignKey:UserID"`
	Ratings              []Rating     `json:"ratings" gorm:"foreignKey:UserID"`
	CreatedAt            time.Time    `json:"created_at"`
	UpdatedAt            time.Time    `json:"updated_at"`
}
