package services

import (
	"errors"
	"madang_api/config"
	"madang_api/models"
)

type FoodService struct{}

// Add a new food item return the food or error if it exist and also check if the food exist for that restaurant
func (s *FoodService) AddFood(food *models.Food) (*models.Food, error) {
	var existingFood models.Food
	//Check if the food exist for that restaurant
	if err := config.DB.Where("restaurant_id = ? AND name = ? ", existingFood.RestaurantID, existingFood.Name).First(&models.Food{}).Error; err == nil {
		return nil, errors.New("food already exists for this restaurant")
	}

	//Add the food item
	if err := config.DB.Create(&food).Error; err != nil {
		return nil, err
	}

	return food, nil
}

// UpdateFood updates an existing food item and returns the updated food item or an error if it fails
func (s *FoodService) UpdateFood(food *models.Food) (*models.Food, error) {
	if err := config.DB.Save(&food).Error; err != nil {
		return nil, err
	}
	return food, nil
}

// Delete a Food by ID
func (s *FoodService) DeleteFood(id uint) error {
	var food models.Food
	if err := config.DB.First(&food, id).Error; err != nil {
		return err
	}
	if err := config.DB.Delete(&food).Error; err != nil {
		return err
	}
	return nil
}

// GetFood retrieves a food item by its ID and returns the food item or an error if it fails
func (s *FoodService) GetFood(id uint) (*models.Food, error) {
	var food models.Food
	if err := config.DB.First(&food, id).Error; err != nil {
		return nil, err
	}
	return &food, nil
}

// GetAllFoods retrieves all food items and returns a slice of food items or an error if it fails
func (s *FoodService) GetAllFoods() ([]models.Food, error) {
	var foods []models.Food
	if err := config.DB.Find(&foods).Error; err != nil {
		return nil, err
	}
	return foods, nil
}

// GetRestaurantFoods retrieves all food items for a specific restaurant and returns a slice of food items or an error if it fails
func (s *FoodService) GetRestaurantFoods(restaurantID uint) ([]models.Food, error) {
	var foods []models.Food
	if err := config.DB.Where("restaurant_id = ?", restaurantID).Find(&foods).Error; err != nil {
		return nil, err
	}
	return foods, nil
}

// implement search for food it should be by name or fhe name of the category the food belongs too
func (s *FoodService) SearchFoods(query string) ([]models.Food, error) {
	var foods []models.Food
	if err := config.DB.Where("name LIKE ? OR category_id IN (SELECT id FROM categories WHERE name LIKE ?)", "%"+query+"%", "%"+query+"%").Find(&foods).Error; err != nil {
		return nil, err
	}
	return foods, nil
}

// get recommended foods for a particular restaurant which is the 5 most recent foods
func (s *FoodService) GetRecommendedFoods(restaurantID uint) ([]models.Food, error) {
	var foods []models.Food
	if err := config.DB.Where("restaurant_id = ?", restaurantID).Order("id desc").Limit(5).Find(&foods).Error; err != nil {
		return nil, err
	}
	return foods, nil
}
