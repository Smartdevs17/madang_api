package services

import (
	"errors"
	"madang_api/config"
	"madang_api/models"
)

type RestaurantService struct{}

// Add a new restaurant
func (s *RestaurantService) AddRestaurant(restaurant *models.Restaurant) (*models.Restaurant, error) {
	// Check if the restaurant already exists
	var existingRestaurant models.Restaurant
	if err := config.DB.Where("name = ?", restaurant.Name).First(&existingRestaurant).Error; err == nil {
		return nil, errors.New("restaurant already exists")
	}
	//verify that the user creating the restaurant has a role of manager and is email verified
	var user models.User
	// if the user is not found, return an error
	if err := config.DB.First(&user, restaurant.UserID).Error; err != nil {
		return nil, errors.New("user not found")
	}
	// if the user is not a manager or the email is not verified, return an error
	if user.Role != "manager" || user.EmailVerified == false {
		return nil, errors.New("user is not authorized to create a restaurant")
	}

	// Add the new restaurant
	if err := config.DB.Create(&restaurant).Error; err != nil {
		return nil, err
	}

	return restaurant, nil
}

// Get a restaurant by ID
func (s *RestaurantService) GetRestaurantByID(id uint) (models.Restaurant, error) {
	var restaurant models.Restaurant
	if err := config.DB.First(&restaurant, id).Error; err != nil {
		return models.Restaurant{}, err
	}
	return restaurant, nil
}

// Get all restaurants
func (s *RestaurantService) GetAllRestaurants() ([]models.Restaurant, error) {
	var restaurants []models.Restaurant
	if err := config.DB.Find(&restaurants).Error; err != nil {
		return nil, err
	}
	return restaurants, nil
}

// Update a restaurant by ID
func (s *RestaurantService) UpdateRestaurant(restaurant *models.Restaurant) (models.Restaurant, error) {
	// Check if the restaurant exists
	var existingRestaurant models.Restaurant
	if err := config.DB.First(&existingRestaurant, restaurant.ID).Error; err != nil {
		return models.Restaurant{}, err
	}

	// Update only the fields that are provided
	if restaurant.Name != "" {
		existingRestaurant.Name = restaurant.Name
	}
	if restaurant.Address != "" {
		existingRestaurant.Address = restaurant.Address
	}
	if restaurant.Location != "" {
		existingRestaurant.Location = restaurant.Location
	}
	if restaurant.Verified != existingRestaurant.Verified {
		existingRestaurant.Verified = restaurant.Verified
	}
	if restaurant.Active != existingRestaurant.Active {
		existingRestaurant.Active = restaurant.Active
	}
	if restaurant.VerfiedAt != "" {
		existingRestaurant.VerfiedAt = restaurant.VerfiedAt
	}
	if restaurant.State != "" {
		existingRestaurant.State = restaurant.State
	}
	if restaurant.Country != "" {
		existingRestaurant.Country = restaurant.Country
	}
	if restaurant.Phone != "" {
		existingRestaurant.Phone = restaurant.Phone
	}
	if restaurant.UserID != 0 {
		// Ensure the user ID is valid
		var user models.User
		if err := config.DB.First(&user, restaurant.UserID).Error; err != nil {
			return models.Restaurant{}, err // Return error if user does not exist
		}
		existingRestaurant.UserID = restaurant.UserID
	}

	// Save the updated restaurant
	if err := config.DB.Save(&existingRestaurant).Error; err != nil {
		return models.Restaurant{}, err
	}

	return existingRestaurant, nil
}

// Delete a restaurant by ID
func (s *RestaurantService) DeleteRestaurant(id uint) error {
	var restaurant models.Restaurant
	if err := config.DB.First(&restaurant, id).Error; err != nil {
		return err
	}
	if err := config.DB.Delete(&restaurant).Error; err != nil {
		return err
	}
	return nil
}

// Get All Verified Restaurants
func (s *RestaurantService) GetAllVerifiedRestaurants() ([]models.Restaurant, error) {
	var restaurants []models.Restaurant
	if err := config.DB.Where("verified = ?", true).Find(&restaurants).Error; err != nil {
		return nil, err
	}
	return restaurants, nil
}

// Get User Restaurants	by ID
func (s *RestaurantService) GetUserRestaurants(id uint) ([]models.Restaurant, error) {
	var restaurants []models.Restaurant
	if err := config.DB.Where("user_id = ?", id).Find(&restaurants).Error; err != nil {
		return nil, err
	}
	return restaurants, nil
}

// search restaurant by name,location,address,state,country
func (s *RestaurantService) SearchRestaurants(query string) ([]models.Restaurant, error) {
	var restaurants []models.Restaurant
	if err := config.DB.Where("name LIKE ? OR location LIKE ? OR address LIKE ? OR state LIKE ? OR country LIKE ?", "%"+query+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%").Find(&restaurants).Error; err != nil {
		return nil, err
	}
	return restaurants, nil
}

// filter restaurant by rating,location
func (s *RestaurantService) FilterRestaurants(state string, country string, location string) ([]models.Restaurant, error) {
	var restaurants []models.Restaurant
	if err := config.DB.Where("state = ? AND location = ? AND country = ? AND ", state, location, country).Find(&restaurants).Error; err != nil {
		return nil, err
	}
	return restaurants, nil
}

// Get All Restaurants with pagination
