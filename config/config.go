package config

import "madang_api/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Restaurant{})
	DB.AutoMigrate(&models.Category{})
	DB.AutoMigrate(&models.Food{})
	DB.AutoMigrate(&models.Table{})
	DB.AutoMigrate(&models.Addon{})
	DB.AutoMigrate(&models.Order{})
	DB.AutoMigrate(&models.Payment{})
	DB.AutoMigrate(&models.Transaction{})
	// DB.AutoMigrate(&models.Rating{})
}
