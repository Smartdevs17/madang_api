package controllers

import (
	"madang_api/models"
	"madang_api/services"
	"madang_api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RestaurantController struct {
	RestaurantService services.RestaurantService
}

type RestaurantControllerInterface interface {
	GetAllRestaurant(c *gin.Context)
	GetRestaurant(c *gin.Context)
	CreateRestaurant(c *gin.Context)
	UpdateRestaurant(c *gin.Context)
	DeleteRestaurant(c *gin.Context)
	SearchRestaurant(c *gin.Context)
	FilterRestaurant(c *gin.Context)
	GetAllVerifiedRestaurants(c *gin.Context)
	GetUserRestaurants(c *gin.Context)
}

// Add a new restaurant
func (ctrl *RestaurantController) CreateRestaurant(c *gin.Context) {
	var body struct {
		Name     string `json:"name"`
		Address  string `json:"address"`
		Location string `json:"location"`
		UserID   uint   `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	if err := utils.ValidateStruct(c, body); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Validation error", err.Error())
		return
	}
	var restaurant models.Restaurant
	restaurant.Name = body.Name
	restaurant.Address = body.Address
	restaurant.Location = body.Location
	restaurant.UserID = body.UserID
	result, err := ctrl.RestaurantService.AddRestaurant(&restaurant)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to register restaurant", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Restaurant created successfully", result)
}

// Get all restaurants
func (ctrl *RestaurantController) GetAllRestaurant(c *gin.Context) {
	restaurants, err := ctrl.RestaurantService.GetAllRestaurants()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch restaurants", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Restaurants fetched successfully", restaurants)
}

// Get a single restaurant
func (ctrl *RestaurantController) GetRestaurant(c *gin.Context) {
	restaurantID, valid := utils.ValidateID(c, "id")
	if !valid {
		return
	}
	restaurant, err := ctrl.RestaurantService.GetRestaurantByID(restaurantID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch restaurant", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Restaurant fetched successfully", restaurant)
}

// Update a restaurant
func (ctrl *RestaurantController) UpdateRestaurant(c *gin.Context) {
	restaurantID, valid := utils.ValidateID(c, "id")
	if !valid {
		return
	}

	var body struct {
		Name      *string `json:"name"`
		Address   *string `json:"address"`
		Location  *string `json:"location"`
		Active    *bool   `json:"active"`
		Verified  *bool   `json:"verified"`
		VerfiedAt *string `json:"verified_at"`
		State     *string `json:"state"`
		Country   *string `json:"country"`
		Phone     *string `json:"phone"`
		UserID    *uint   `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	// Fetch existing restaurant
	existingRestaurant, err := ctrl.RestaurantService.GetRestaurantByID(restaurantID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Restaurant not found", err.Error())
		return
	}

	// Update fields if they are provided
	if body.Name != nil {
		existingRestaurant.Name = *body.Name
	}
	if body.Address != nil {
		existingRestaurant.Address = *body.Address
	}
	if body.Location != nil {
		existingRestaurant.Location = *body.Location
	}
	if body.Active != nil {
		existingRestaurant.Active = *body.Active
	}
	if body.Verified != nil {
		existingRestaurant.Verified = *body.Verified
	}
	if body.VerfiedAt != nil {
		existingRestaurant.VerfiedAt = *body.VerfiedAt
	}
	if body.State != nil {
		existingRestaurant.State = *body.State
	}
	if body.Country != nil {
		existingRestaurant.Country = *body.Country
	}
	if body.Phone != nil {
		existingRestaurant.Phone = *body.Phone
	}
	if body.UserID != nil {
		existingRestaurant.UserID = *body.UserID
	}

	updatedRestaurant, err := ctrl.RestaurantService.UpdateRestaurant(&existingRestaurant)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update restaurant", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Restaurant updated successfully", updatedRestaurant)
}

// Delete a restaurant
func (ctrl *RestaurantController) DeleteRestaurant(c *gin.Context) {
	restaurantID, valid := utils.ValidateID(c, "id")
	if !valid {
		return
	}

	err := ctrl.RestaurantService.DeleteRestaurant(restaurantID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to delete restaurant", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Restaurant deleted successfully", nil)
}

// Search for restaurants
func (ctrl *RestaurantController) SearchRestaurant(c *gin.Context) {
	query := c.Query("q")
	restaurants, err := ctrl.RestaurantService.SearchRestaurants(query)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to search restaurants", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Restaurants searched successfully", restaurants)
}

// Filter restaurants
func (ctrl *RestaurantController) FilterRestaurant(c *gin.Context) {
	var body struct {
		State    string `json:"state" binding:"required"`
		Country  string `json:"country" binding:"required"`
		Location string `json:"location" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	if err := utils.ValidateStruct(c, body); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Validation error", err.Error())
		return
	}

	restaurants, err := ctrl.RestaurantService.FilterRestaurants(body.State, body.Country, body.Location)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to filter restaurants", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Restaurants filtered successfully", restaurants)
}

// Get all verified restaurants
func (ctrl *RestaurantController) GetAllVerifiedRestaurants(c *gin.Context) {
	restaurants, err := ctrl.RestaurantService.GetAllVerifiedRestaurants()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch verified restaurants", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Verified restaurants fetched successfully", restaurants)
}

// Get user restaurants
func (ctrl *RestaurantController) GetUserRestaurants(c *gin.Context) {
	userID, valid := utils.ValidateID(c, "user_id")
	if !valid {
		return
	}

	restaurants, err := ctrl.RestaurantService.GetUserRestaurants(userID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch user restaurants", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "User restaurants fetched successfully", restaurants)
}
