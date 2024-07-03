package services

import (
	"errors"
	"madang_api/config"
	"madang_api/models"
)

type CategoryService struct{}

func (s *CategoryService) GetCategories() ([]models.Category, error) {
	var categories []models.Category
	result := config.DB.Find(&categories)
	return categories, result.Error
}

func (s *CategoryService) GetCategory(id uint) (models.Category, error) {
	var category models.Category
	result := config.DB.First(&category, id)
	return category, result.Error
}

func (s *CategoryService) CreateCategory(category *models.Category) (*models.Category, error) {
	var existingCategory models.Category
	if err := config.DB.Where("name = ? AND restaurant_id = ?", category.Name, category.RestaurantID).First(&existingCategory).Error; err == nil {
		return nil, errors.New("a category with the same name already exists")
	}

	result := config.DB.Create(category)
	if result.Error != nil {
		return nil, result.Error
	}

	return category, nil
}

func (s *CategoryService) UpdateCategory(category *models.Category) (*models.Category, error) {
	result := config.DB.Save(category)
	if result.Error != nil {
		return nil, result.Error
	}

	return category, nil
}

func (s *CategoryService) DeleteCategory(id uint) error {
	result := config.DB.Delete(&models.Category{}, id)
	return result.Error
}

func (s *CategoryService) GetRestaurantCategories(restaurantID uint) ([]models.Category, error) {
	var categories []models.Category
	result := config.DB.Where("restaurant_id = ?", restaurantID).Find(&categories)
	return categories, result.Error
}
