package controllers

import (
	"madang_api/models"
	"madang_api/services"
	"madang_api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FoodControllerImpl implements the FoodController interface
type FoodController struct {
	FoodService services.FoodService
}

// FoodController defines the methods for handling food related operations
type FoodControllerInterface interface {
	AddFood(ctx *gin.Context)
	UpdateFood(ctx *gin.Context)
	DeleteFood(ctx *gin.Context)
	GetFood(ctx *gin.Context)
	GetAllFoods(ctx *gin.Context)
	GetRestaurantFoods(ctx *gin.Context)
}

// AddFood handles the addition of a new food item
func (ctrl *FoodController) AddFood(c *gin.Context) {
	var body struct {
		Name         string  `json:"name"`
		Description  string  `json:"description"`
		Image        string  `json:"image"`
		Price        float64 `json:"price"`
		RestaurantID uint    `json:"restaurant_id"`
		CategoryId   uint    `json:"category_id"`
	}

	// Validate the request body
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	if err := utils.ValidateStruct(c, body); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Validation error", err.Error())
		return
	}
	var food models.Food
	food.Name = body.Name
	food.Description = body.Description
	food.Image = body.Image
	food.Price = body.Price
	food.RestaurantID = body.RestaurantID
	food.CategoryId = body.CategoryId

	// Call the AddFood service
	newFood, err := ctrl.FoodService.AddFood(&food)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to add food", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Food added successfully", newFood)

}

// UpdateFood handles the update of an existing food item
func (f *FoodController) UpdateFood(c *gin.Context) {
	foodId, valid := utils.ValidateID(c, "id")
	if !valid {
		return
	}
	var body struct {
		Name         string  `json:"name"`
		Description  string  `json:"description"`
		Image        string  `json:"image"`
		Price        float64 `json:"price"`
		CategoryId   uint    `json:"category_id"`
		RestaurantId uint    `json:"restaurant_id"`
	}

	// Validate the request body
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	// Check if the food item exists
	existingFood, err := f.FoodService.GetFood(foodId)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Food not found", err.Error())
		return
	}

	if len(body.Name) == 0 {
		body.Name = existingFood.Name
	}

	if len(body.Description) == 0 {
		body.Description = existingFood.Description
	}

	if len(body.Image) == 0 {
		body.Image = existingFood.Image
	}

	if body.Price == 0 {
		body.Price = existingFood.Price
	}

	if body.CategoryId == 0 {
		body.CategoryId = existingFood.CategoryId
	}

	if body.RestaurantId == 0 {
		body.RestaurantId = existingFood.RestaurantID
	}

	if err := utils.ValidateStruct(c, body); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Validation error", err.Error())
		return
	}

	var food models.Food
	food.Name = body.Name
	food.Description = body.Description
	food.Image = body.Image
	food.Price = body.Price
	food.CategoryId = body.CategoryId
	food.RestaurantID = body.RestaurantId

	// Call the UpdateFood service
	updatedFood, err := f.FoodService.UpdateFood(&food)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update food", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Food updated successfully", updatedFood)
}

// DeleteFood handles the deletion of a food item
func (f *FoodController) DeleteFood(c *gin.Context) {
	foodId, valid := utils.ValidateID(c, "id")
	if !valid {
		return
	}

	_, err := f.FoodService.GetFood(foodId)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Food not found", err.Error())
		return
	}

	// Call the DeleteFood service
	err = f.FoodService.DeleteFood(foodId)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete food", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Food deleted successfully", nil)
}

// GetFood retrieves a specific food item
func (f *FoodController) GetFood(c *gin.Context) {
	foodId, valid := utils.ValidateID(c, "id")
	if !valid {
		return
	}

	// Call the GetFood service
	food, err := f.FoodService.GetFood(foodId)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve food", err.Error())
		return
	}
	// Return the food item
	utils.SuccessResponse(c, http.StatusOK, "Food retrieved successfully", food)
}

// GetAllFoods retrieves all food items
func (f *FoodController) GetAllFoods(c *gin.Context) {
	// Call the GetAllFoods service
	foods, err := f.FoodService.GetAllFoods()
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve foods", err.Error())
		return
	}

	// Return the list of foods
	utils.SuccessResponse(c, http.StatusOK, "Foods retrieved successfully", foods)
}

// GetRestaurantFoods retrieves all the foods of a particular restaurant
func (f *FoodController) GetRestaurantFoods(c *gin.Context) {
	restaurantId, valid := utils.ValidateID(c, "id")
	if !valid {
		return
	}

	// Call the GetRestaurantFoods service
	foods, err := f.FoodService.GetRestaurantFoods(restaurantId)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve foods", err.Error())
		return
	}

	// Return the list of foods
	utils.SuccessResponse(c, http.StatusOK, "Foods retrieved successfully", foods)
}

// SearchFood
func (f *FoodController) SearchFood(c *gin.Context) {
	// Get the search query parameter
	query := c.Query("q")

	// Call the SearchFood service
	foods, err := f.FoodService.SearchFoods(query)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to search for foods", err.Error())
		return
	}

	// Return the list of foods
	utils.SuccessResponse(c, http.StatusOK, "Foods retrieved successfully", foods)
}
