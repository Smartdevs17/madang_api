package services

import (
	"madang_api/config"
	"madang_api/models"
)

type OrderService struct{}

// Add a new order item return the order or error if it exist and also check if the order exist for that restaurant
func (s *OrderService) AddOrder(order *models.Order) (*models.Order, error) {
	//Add the order item
	if err := config.DB.Create(&order).Error; err != nil {
		return nil, err
	}

	return order, nil
}

// UpdateOrder updates an existing order item and returns the updated order item or an error if it fails
func (s *OrderService) UpdateOrder(order *models.Order) (*models.Order, error) {
	if err := config.DB.Save(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

// Delete a Order by ID
func (s *OrderService) DeleteOrder(id uint) error {
	var order models.Order
	if err := config.DB.First(&order, id).Error; err != nil {
		return err
	}
	if err := config.DB.Delete(&order).Error; err != nil {
		return err
	}
	return nil
}

// GetOrder retrieves a order item by its ID and returns the order item or an error if it fails
func (s *OrderService) GetOrder(id uint) (*models.Order, error) {
	var order models.Order
	if err := config.DB.First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

// GetAllOrders retrieves all order items and returns a slice of order items or an error if it fails
func (s *OrderService) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	if err := config.DB.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// GetRestaurantOrders retrieves all order items for a specific restaurant and returns a slice of order items or an error if it fails
func (s *OrderService) GetRestaurantOrders(restaurantID uint) ([]models.Order, error) {
	var orders []models.Order
	if err := config.DB.Where("restaurant_id = ?", restaurantID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// implement search for order it should be by name or fhe name of the category the order belongs too
func (s *OrderService) SearchOrders(query string) ([]models.Order, error) {
	var orders []models.Order
	if err := config.DB.Where("name LIKE ? OR category_id IN (SELECT id FROM categories WHERE name LIKE ?)", "%"+query+"%", "%"+query+"%").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
