package controllers

import (
	"madang_api/models"
	"madang_api/services"
	"madang_api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddonControllerImpl implements the AddonController interface
type AddonController struct {
	AddonService services.AddonService
}

// AddonController defines the methods for handling addon related operations
type AddonControllerInterface interface {
	AddAddon(ctx *gin.Context)
	UpdateAddon(ctx *gin.Context)
	DeleteAddon(ctx *gin.Context)
	GetAddon(ctx *gin.Context)
	GetAllAddons(ctx *gin.Context)
	GetRestaurantAddons(ctx *gin.Context)
}

// AddAddon handles the addition of a new addon item
func (ctrl *AddonController) AddAddon(c *gin.Context) {
	var body struct {
		Name         string  `json:"name"`
		Type         string  `json:"type"`
		Price        float64 `json:"price"`
		RestaurantID uint    `json:"restaurant_id"`
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
	var addon models.Addon
	addon.Name = body.Name
	addon.Type = body.Type
	addon.Price = body.Price
	addon.RestaurantID = body.RestaurantID

	// Call the AddAddon service
	newAddon, err := ctrl.AddonService.AddAddon(&addon)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to add addon", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Addon added successfully", newAddon)

}

// UpdateAddon handles the update of an existing addon item
func (f *AddonController) UpdateAddon(c *gin.Context) {
	addonId, valid := utils.ValidateID(c, "id")
	if !valid {
		return
	}
	var body struct {
		Name         string  `json:"name"`
		Type         string  `json:"type"`
		Price        float64 `json:"price"`
		RestaurantID uint    `json:"restaurant_id"`
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

	// Check if the addon item exists
	_, err := f.AddonService.GetAddon(addonId)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Addon not found", err.Error())
		return
	}

	var addon models.Addon
	addon.Name = body.Name
	addon.Type = body.Type
	addon.Price = body.Price

	// Call the UpdateAddon service
	updatedAddon, err := f.AddonService.UpdateAddon(&addon)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update addon", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Addon updated successfully", updatedAddon)
}

// DeleteAddon handles the deletion of a addon item
func (f *AddonController) DeleteAddon(c *gin.Context) {
	addonId, valid := utils.ValidateID(c, "id")
	if !valid {
		return
	}

	// Call the DeleteAddon service
	err := f.AddonService.DeleteAddon(addonId)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete addon", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Addon deleted successfully", nil)
}

// GetAddon retrieves a specific addon item
func (f *AddonController) GetAddon(c *gin.Context) {
	addonId, valid := utils.ValidateID(c, "id")
	if !valid {
		return
	}

	// Call the GetAddon service
	addon, err := f.AddonService.GetAddon(addonId)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve addon", err.Error())
		return
	}
	// Return the addon item
	utils.SuccessResponse(c, http.StatusOK, "Addon retrieved successfully", addon)
}

// GetAllAddons retrieves all addon items
func (f *AddonController) GetAllAddons(c *gin.Context) {
	// Call the GetAllAddons service
	addons, err := f.AddonService.GetAllAddons()
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve addons", err.Error())
		return
	}

	// Return the list of addons
	utils.SuccessResponse(c, http.StatusOK, "Addons retrieved successfully", addons)
}

// GetRestaurantAddons retrieves all the addons of a particular restaurant
func (f *AddonController) GetRestaurantAddons(c *gin.Context) {
	restaurantId, valid := utils.ValidateID(c, "restaurant_id")
	if !valid {
		return
	}

	// Call the GetRestaurantAddons service
	addons, err := f.AddonService.GetRestaurantAddons(restaurantId)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve addons", err.Error())
		return
	}

	// Return the list of addons
	utils.SuccessResponse(c, http.StatusOK, "Addons retrieved successfully", addons)
}

// SearchAddon
func (f *AddonController) SearchAddon(c *gin.Context) {
	// Get the search query parameter
	query := c.Query("q")

	// Call the SearchAddon service
	addons, err := f.AddonService.SearchAddons(query)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to search for addons", err.Error())
		return
	}

	// Return the list of addons
	utils.SuccessResponse(c, http.StatusOK, "Addons retrieved successfully", addons)
}
