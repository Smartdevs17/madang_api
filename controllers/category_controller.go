package controllers

import (
	"madang_api/models"
	"madang_api/services"
	"madang_api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	CategoryService services.CategoryService
}

type CategoryControllerInterface interface {
	GetAllCategories(c *gin.Context)
	GetCategory(c *gin.Context)
	CreateCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
	GetRestaurantCategories(c *gin.Context)
}

func (ctrl *CategoryController) GetAllCategories(c *gin.Context) {
	categories, err := ctrl.CategoryService.GetCategories()
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve categories", err.Error())
		return
	}

	// Return the list of foods
	utils.SuccessResponse(c, http.StatusOK, "Categories retrieved successfully", categories)
}

func (ctrl *CategoryController) GetCategory(c *gin.Context) {
	// Get the category ID from the request parameters
	categoryID, valid := utils.ValidateID(c, "id")
	if !valid {
		return
	}

	// Call the service to retrieve the category
	category, err := ctrl.CategoryService.GetCategory(categoryID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve category", err.Error())
		return
	}

	// Return the category
	utils.SuccessResponse(c, http.StatusOK, "Category retrieved successfully", category)
}

func (ctrl *CategoryController) CreateCategory(c *gin.Context) {
	// Bind the request body to a Category struct
	var body struct {
		Name         string `json:"name"`
		Type         string `json:"type"`
		RestaurantID uint   `json:"restaurant_id"`
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
	var category models.Category
	category.Name = body.Name
	category.Type = body.Type
	category.RestaurantID = body.RestaurantID

	// Call the service to create the category
	newCategory, err := ctrl.CategoryService.CreateCategory(&category)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create category", err.Error())
		return
	}

	// Return the created category
	utils.SuccessResponse(c, http.StatusCreated, "Category created successfully", newCategory)
}

func (ctrl *CategoryController) UpdateCategory(c *gin.Context) {
	// Get the category ID from the request parameters
	categoryID, valid := utils.ValidateID(c, "id")
	if !valid {
		return
	}

	// Bind the request body to a Category struct
	var body struct {
		Name         string `json:"name"`
		Type         string `json:"type"`
		RestaurantID uint   `json:"restaurant_id"`
	}

	// Validate the request body
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}
	//check fig the catergory exists
	_, err := ctrl.CategoryService.GetCategory(categoryID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Category not found", err.Error())
		return
	}
	var category models.Category
	category.ID = categoryID
	category.Name = body.Name
	category.Type = body.Type
	category.RestaurantID = body.RestaurantID

	// Call the service to update the category
	updatedCategory, err := ctrl.CategoryService.UpdateCategory(&category)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update category", err.Error())
		return
	}

	// Return the updated category
	utils.SuccessResponse(c, http.StatusOK, "Category updated successfully", updatedCategory)
}

func (ctrl *CategoryController) DeleteCategory(c *gin.Context) {
	// Get the category ID from the request parameters
	categoryID, valid := utils.ValidateID(c, "id")
	if !valid {
		return
	}

	//check if the category exists
	_, errExist := ctrl.CategoryService.GetCategory(categoryID)
	// Handle error
	if errExist != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Category not found", errExist.Error())
		return
	}

	// Call the service to delete the category
	err := ctrl.CategoryService.DeleteCategory(categoryID)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete category", err.Error())
		return
	}

	// Return a success response
	utils.SuccessResponse(c, http.StatusOK, "Category deleted successfully", nil)
}

func (ctrl *CategoryController) GetRestaurantCategories(c *gin.Context) {
	// Get the restaurant ID from the request parameters
	restaurantID, valid := utils.ValidateID(c, "restaurant_id")
	if !valid {
		return
	}

	// Call the service to retrieve the categories for the restaurant
	categories, err := ctrl.CategoryService.GetRestaurantCategories(restaurantID)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve categories", err.Error())
		return
	}

	// Return the list of categories
	utils.SuccessResponse(c, http.StatusOK, "Categories retrieved successfully", categories)
}
