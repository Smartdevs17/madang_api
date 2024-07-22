package controllers

import (
	"madang_api/models"
	"madang_api/services"
	"madang_api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TableControllerImpl implements the TableController interface
type TableController struct {
	TableService services.TableService
}

// TableController defines the methods for handling table related operations
type TableControllerInterface interface {
	AddTable(ctx *gin.Context)
	UpdateTable(ctx *gin.Context)
	DeleteTable(ctx *gin.Context)
	GetTable(ctx *gin.Context)
	GetAllTables(ctx *gin.Context)
	GetRestaurantTables(ctx *gin.Context)
}

// AddTable handles the addition of a new table item
func (ctrl *TableController) AddTable(c *gin.Context) {
	var body struct {
		Name         string  `json:"name"`
		Number       float64 `json:"number"`
		Capacity     float64 `json:"capacity"`
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
	var table models.Table
	table.Name = body.Name
	table.Capacity = int(body.Capacity)
	table.Number = int(body.Number)
	table.Image = body.Image
	table.Price = body.Price
	table.RestaurantID = body.RestaurantID
	table.CategoryId = body.CategoryId

	// Call the AddTable service
	newTable, err := ctrl.TableService.AddTable(&table)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to add table", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Table added successfully", newTable)

}

// UpdateTable handles the update of an existing table item
func (f *TableController) UpdateTable(c *gin.Context) {
	tableId, valid := utils.ValidateID(c, "id")
	if !valid {
		return
	}
	var body struct {
		Name         string  `json:"name"`
		Number       int     `json:"number"`
		Capacity     int     `json:"capacity"`
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

	// Check if the table item exists
	existingTable, err := f.TableService.GetTable(tableId)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Table not found", err.Error())
		return
	}
	// Update the table fields
	if body.Name == "" {
		body.Name = existingTable.Name
	}
	if body.Number == 0 {
		body.Number = existingTable.Number
	}
	if body.Capacity == 0 {
		body.Capacity = existingTable.Capacity
	}
	if body.Image == "" {
		body.Image = existingTable.Image
	}
	if body.Price == 0 {
		body.Price = existingTable.Price
	}
	if body.RestaurantID == 0 {
		body.RestaurantID = existingTable.RestaurantID
	}
	if body.CategoryId == 0 {
		body.CategoryId = existingTable.CategoryId
	}

	if err := utils.ValidateStruct(c, body); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Validation error", err.Error())
		return
	}

	var table models.Table
	table.Name = body.Name
	table.Number = body.Number
	table.Capacity = body.Capacity
	table.Image = body.Image
	table.Price = body.Price
	table.RestaurantID = body.RestaurantID
	table.CategoryId = body.CategoryId

	// Call the UpdateTable service
	updatedTable, err := f.TableService.UpdateTable(&table)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update table", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Table updated successfully", updatedTable)
}

// DeleteTable handles the deletion of a table item
func (f *TableController) DeleteTable(c *gin.Context) {
	tableId, valid := utils.ValidateID(c, "id")
	if !valid {
		return
	}

	_, err := f.TableService.GetTable(tableId)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Table not found", err.Error())
		return
	}

	// Call the DeleteTable service
	err = f.TableService.DeleteTable(tableId)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete table", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Table deleted successfully", nil)
}

// GetTable retrieves a specific table item
func (f *TableController) GetTable(c *gin.Context) {
	tableId, valid := utils.ValidateID(c, "id")
	if !valid {
		return
	}

	// Call the GetTable service
	table, err := f.TableService.GetTable(tableId)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Failed to retrieve table", err.Error())
		return
	}
	// Return the table item
	utils.SuccessResponse(c, http.StatusOK, "Table retrieved successfully", table)
}

// GetAllTables retrieves all table items
func (f *TableController) GetAllTables(c *gin.Context) {
	// Call the GetAllTables service
	tables, err := f.TableService.GetAllTables()
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve tables", err.Error())
		return
	}

	// Return the list of tables
	utils.SuccessResponse(c, http.StatusOK, "Tables retrieved successfully", tables)
}

// GetRestaurantTables retrieves all the tables of a particular restaurant
func (f *TableController) GetRestaurantTables(c *gin.Context) {
	restaurantId, valid := utils.ValidateID(c, "id")
	if !valid {
		return
	}

	// Call the GetRestaurantTables service
	tables, err := f.TableService.GetRestaurantTables(restaurantId)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve tables", err.Error())
		return
	}

	// Return the list of tables
	utils.SuccessResponse(c, http.StatusOK, "Tables retrieved successfully", tables)
}

// SearchTable
func (f *TableController) SearchTable(c *gin.Context) {
	// Get the search query parameter
	query := c.Query("q")

	// Call the SearchTable service
	tables, err := f.TableService.SearchTables(query)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to search for tables", err.Error())
		return
	}

	// Return the list of tables
	utils.SuccessResponse(c, http.StatusOK, "Tables retrieved successfully", tables)
}
