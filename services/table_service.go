package services

import (
	"errors"
	"madang_api/config"
	"madang_api/models"
)

type TableService struct{}

// Add a new table item return the table or error if it exist and also check if the table exist for that restaurant
func (s *TableService) AddTable(table *models.Table) (*models.Table, error) {
	var existingTable models.Table
	//Check if the table exist for that restaurant
	if err := config.DB.Where("restaurant_id = ? AND name = ? ", existingTable.RestaurantID, existingTable.Name).First(&models.Table{}).Error; err == nil {
		return nil, errors.New("table already exists for this restaurant")
	}

	//Add the table item
	if err := config.DB.Create(&table).Error; err != nil {
		return nil, err
	}

	return table, nil
}

// UpdateTable updates an existing table item and returns the updated table item or an error if it fails
func (s *TableService) UpdateTable(table *models.Table) (*models.Table, error) {
	if err := config.DB.Save(&table).Error; err != nil {
		return nil, err
	}
	return table, nil
}

// Delete a Table by ID
func (s *TableService) DeleteTable(id uint) error {
	var table models.Table
	if err := config.DB.First(&table, id).Error; err != nil {
		return err
	}
	if err := config.DB.Delete(&table).Error; err != nil {
		return err
	}
	return nil
}

// GetTable retrieves a table item by its ID and returns the table item or an error if it fails
func (s *TableService) GetTable(id uint) (*models.Table, error) {
	var table models.Table
	if err := config.DB.First(&table, id).Error; err != nil {
		return nil, err
	}
	return &table, nil
}

// GetAllTables retrieves all table items and returns a slice of table items or an error if it fails
func (s *TableService) GetAllTables() ([]models.Table, error) {
	var tables []models.Table
	if err := config.DB.Find(&tables).Error; err != nil {
		return nil, err
	}
	return tables, nil
}

// GetRestaurantTables retrieves all table items for a specific restaurant and returns a slice of table items or an error if it fails
func (s *TableService) GetRestaurantTables(restaurantID uint) ([]models.Table, error) {
	var tables []models.Table
	if err := config.DB.Where("restaurant_id = ?", restaurantID).Find(&tables).Error; err != nil {
		return nil, err
	}
	return tables, nil
}

// implement search for table it should be by name or fhe name of the category the table belongs too
func (s *TableService) SearchTables(query string) ([]models.Table, error) {
	var tables []models.Table
	if err := config.DB.Where("name LIKE ? OR category_id IN (SELECT id FROM categories WHERE name LIKE ?)", "%"+query+"%", "%"+query+"%").Find(&tables).Error; err != nil {
		return nil, err
	}
	return tables, nil
}

// get recommended tables for a particular restaurant which is the 5 most recent tables
func (s *TableService) GetRecommendedTables(restaurantID uint) ([]models.Table, error) {
	var tables []models.Table
	if err := config.DB.Where("restaurant_id = ?", restaurantID).Order("id desc").Limit(5).Find(&tables).Error; err != nil {
		return nil, err
	}
	return tables, nil
}
