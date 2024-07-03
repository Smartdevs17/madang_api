package services

import (
	"errors"
	"madang_api/config"
	"madang_api/models"
)

type AddonService struct{}

// Add a new addon item return the addon or error if it exist and also check if the addon exist for that restaurant
func (s *AddonService) AddAddon(addon *models.Addon) (*models.Addon, error) {
	var existingAddon models.Addon
	//Check if the addon exist for that restaurant
	if err := config.DB.Where("restaurant_id = ? AND name = ? ", existingAddon.RestaurantID, existingAddon.Name).First(&models.Addon{}).Error; err == nil {
		return nil, errors.New("addon already exists for this restaurant")
	}

	//Add the addon item
	if err := config.DB.Create(&addon).Error; err != nil {
		return nil, err
	}

	return addon, nil
}

// UpdateAddon updates an existing addon item and returns the updated addon item or an error if it fails
func (s *AddonService) UpdateAddon(addon *models.Addon) (*models.Addon, error) {
	if err := config.DB.Save(&addon).Error; err != nil {
		return nil, err
	}
	return addon, nil
}

// Delete a Addon by ID
func (s *AddonService) DeleteAddon(id uint) error {
	var addon models.Addon
	if err := config.DB.First(&addon, id).Error; err != nil {
		return err
	}
	if err := config.DB.Delete(&addon).Error; err != nil {
		return err
	}
	return nil
}

// GetAddon retrieves a addon item by its ID and returns the addon item or an error if it fails
func (s *AddonService) GetAddon(id uint) (*models.Addon, error) {
	var addon models.Addon
	if err := config.DB.First(&addon, id).Error; err != nil {
		return nil, err
	}
	return &addon, nil
}

// GetAllAddons retrieves all addon items and returns a slice of addon items or an error if it fails
func (s *AddonService) GetAllAddons() ([]models.Addon, error) {
	var addons []models.Addon
	if err := config.DB.Find(&addons).Error; err != nil {
		return nil, err
	}
	return addons, nil
}

// GetRestaurantAddons retrieves all addon items for a specific restaurant and returns a slice of addon items or an error if it fails
func (s *AddonService) GetRestaurantAddons(restaurantID uint) ([]models.Addon, error) {
	var addons []models.Addon
	if err := config.DB.Where("restaurant_id = ?", restaurantID).Find(&addons).Error; err != nil {
		return nil, err
	}
	return addons, nil
}

// implement search for addon it should be by name or fhe name of the category the addon belongs too
func (s *AddonService) SearchAddons(query string) ([]models.Addon, error) {
	var addons []models.Addon
	if err := config.DB.Where("name LIKE ? OR category_id IN (SELECT id FROM categories WHERE name LIKE ?)", "%"+query+"%", "%"+query+"%").Find(&addons).Error; err != nil {
		return nil, err
	}
	return addons, nil
}
