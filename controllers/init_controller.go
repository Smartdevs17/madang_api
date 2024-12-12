package controllers

import (
	"madang_api/models"
	"madang_api/services"
	"madang_api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InitController struct {
	UserService     services.UserService
	CategoryService services.CategoryService
	FoodService     services.FoodService
	TableService    services.TableService
	OrderService    services.OrderService
}

// InitController defines the methods for handling initial items been loaded
type InitControllerInterface interface {
	LoadData(ctx *gin.Context)
}

// Get Initialize items
func (f *InitController) LoadData(c *gin.Context) {
	// Get the authenticated user from the context
	loggedInUser, exists := c.Get("user")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "bad request", "User not found")
		return
	}

	// Extract user ID from the user object
	userId := loggedInUser.(models.User).ID

	// Declare a variable to store the items
	var items struct {
		User              UserResponse      `json:"user"`
		Categories        []models.Category `json:"categories"`
		Foods             []models.Food     `json:"foods"`
		Tables            []models.Table    `json:"tables"`
		RecommendedFoods  []models.Food     `json:"recommended_foods"`
		RecommendedTables []models.Table    `json:"recommended_tables"`
		Orders            []models.Order    `json:"orders"`
	}

	//Get the User Details
	user, err := f.UserService.GetUserByID(userId)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve user", err.Error())
		return
	}

	userDetails := UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Avatar:    user.Avatar,
		Role:      user.Role,
		Active:    user.Active,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	// Assign the retrieved user to the items struct
	items.User = userDetails

	//call the Restaurant Category service
	categories, err := f.CategoryService.GetRestaurantCategories(4)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve categories", err.Error())
	}
	items.Categories = categories

	// Call the Restaurant Food service
	foods, err := f.FoodService.GetRestaurantFoods(4)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve foods", err.Error())
		return
	}

	// Assign the retrieved foods to the items struct
	items.Foods = foods

	//call recommended foods service
	recommendedFoods, err := f.FoodService.GetRecommendedFoods(4)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve recommended foods", err.Error())
	}
	// Assign the retrieved foods to the items struct
	items.RecommendedFoods = recommendedFoods

	// Call the Table service
	tables, err := f.TableService.GetRestaurantTables(4) // Declare 'tables' here
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve tables", err.Error())
		return
	}

	// Assign the retrieved tables to the items struct
	items.Tables = tables

	//Call the Recommended Table service
	recommendedTables, err := f.TableService.GetRecommendedTables(4)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve recommended tables", err.Error())
	}
	items.RecommendedTables = recommendedTables

	// Call the Order service
	orders, err := f.OrderService.GetUserOrders(userId)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve orders", err.Error())
	}
	items.Orders = orders

	// Return the items
	utils.SuccessResponse(c, http.StatusOK, "Foods and tables retrieved successfully", items)
}
