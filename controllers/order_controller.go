package controllers

import (
	"madang_api/models"
	"madang_api/services"
	"madang_api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// OrderControllerImpl implements the OrderController interface
type OrderController struct {
	OrderService services.OrderService
}

// OrderController defines the methods for handling order related operations
type OrderControllerInterface interface {
	AddOrder(ctx *gin.Context)
	UpdateOrder(ctx *gin.Context)
	DeleteOrder(ctx *gin.Context)
	GetOrder(ctx *gin.Context)
	GetAllOrders(ctx *gin.Context)
	GetRestaurantOrders(ctx *gin.Context)
}

// AddOrder handles the addition of a new order item
func (ctrl *OrderController) AddOrder(c *gin.Context) {
	var body struct {
		UserID       uint `json:"user_id"`
		TableID      uint `json:"table_id"`
		RestaurantID uint `json:"restaurant_id"`
		Foods        []struct {
			ID       uint `json:"id"`
			Quantity int  `json:"quantity"`
		} `json:"foods"`
		Tables []struct {
			TableID uint `json:"table_id"`
		} `json:"tables"`
		Addons []struct {
			ID       uint `json:"id"`
			Quantity int  `json:"quantity"`
		} `json:"addons"`
		TotalPrice   float64 `json:"total_price"`
		SpecialNotes string  `json:"special_notes"`
		Status       string  `json:"status"`
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

	var order models.Order
	order.UserID = body.UserID
	order.TableID = &body.TableID
	order.RestaurantID = body.RestaurantID
	order.Status = body.Status

	// Convert body.Foods to []models.FoodOrder
	var foodOrders []models.FoodOrder
	for _, food := range body.Foods {
		foodOrder := models.FoodOrder{
			FoodID:   food.ID,
			Quantity: food.Quantity,
		}
		foodOrders = append(foodOrders, foodOrder)
	}
	order.FoodOrders = foodOrders

	// Convert body.Tables to []models.TableOrder
	var tableOrders []models.TableOrder
	for _, table := range body.Tables {
		tableOrder := models.TableOrder{
			TableID: table.TableID,
		}
		tableOrders = append(tableOrders, tableOrder)
	}
	order.TableOrders = tableOrders

	// Convert body.Addons to []models.AddonOrder
	var addonOrders []models.AddonOrder
	for _, addon := range body.Addons {
		addonOrder := models.AddonOrder{
			Addon:    models.Addon{ID: addon.ID},
			Quantity: addon.Quantity,
		}
		addonOrders = append(addonOrders, addonOrder)
	}
	order.AddonOrders = addonOrders

	order.TotalPrice = body.TotalPrice
	order.SpecialNotes = body.SpecialNotes

	// Call the AddOrder service
	newOrder, err := ctrl.OrderService.AddOrder(&order)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to add order", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Order added successfully", newOrder)
}

// UpdateOrder handles the update of an existing order item
func (f *OrderController) UpdateOrder(c *gin.Context) {
	orderId, valid := utils.ValidateID(c, "id")
	if !valid {
		return
	}
	var body struct {
		UserID       uint `json:"user_id"`
		TableID      uint `json:"table_id"`
		RestaurantID uint `json:"restaurant_id"`
		Foods        []struct {
			ID       uint `json:"id"`
			Quantity int  `json:"quantity"`
		} `json:"foods"`
		Tables []struct {
			TableID uint `json:"table_id"`
		} `json:"tables"`
		Addons []struct {
			ID       uint `json:"id"`
			Quantity int  `json:"quantity"`
		} `json:"addons"`
		TotalPrice   float64 `json:"total_price"`
		SpecialNotes string  `json:"special_notes"`
		Status       string  `json:"status"`
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

	// Check if the order item exists
	_, err := f.OrderService.GetOrder(orderId)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Order not found", err.Error())
		return
	}

	var order models.Order
	order.UserID = body.UserID
	order.TableID = &body.TableID
	order.RestaurantID = body.RestaurantID
	order.Status = body.Status

	// Convert body.Foods to []models.FoodOrder
	var foodOrders []models.FoodOrder
	for _, food := range body.Foods {
		foodOrders = append(foodOrders, models.FoodOrder{
			FoodID:   food.ID,
			Quantity: food.Quantity,
		})
	}
	order.FoodOrders = foodOrders

	// Convert body.TableOrders to []models.TableOrder
	var tableOrders []models.TableOrder
	for _, table := range body.Tables {
		tableOrders = append(tableOrders, models.TableOrder{
			TableID: table.TableID,
		})
	}
	order.TableOrders = tableOrders

	// Convert body.Addons to []models.AddonOrder
	var addonOrders []models.AddonOrder
	for _, addon := range body.Addons {
		addonOrders = append(addonOrders, models.AddonOrder{
			Addon:    models.Addon{ID: addon.ID},
			Quantity: addon.Quantity,
		})
	}
	order.AddonOrders = addonOrders

	order.TotalPrice = body.TotalPrice
	order.SpecialNotes = body.SpecialNotes

	// Call the UpdateOrder service
	updatedOrder, err := f.OrderService.UpdateOrder(&order)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update order", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Order updated successfully", updatedOrder)
}

// DeleteOrder handles the deletion of a order item
func (f *OrderController) DeleteOrder(c *gin.Context) {
	orderId, valid := utils.ValidateID(c, "id")
	if !valid {
		return
	}

	// Call the DeleteOrder service
	err := f.OrderService.DeleteOrder(orderId)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete order", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Order deleted successfully", nil)
}

// GetOrder retrieves a specific order item
func (f *OrderController) GetOrder(c *gin.Context) {
	orderId, valid := utils.ValidateID(c, "id")
	if !valid {
		return
	}

	// Call the GetOrder service
	order, err := f.OrderService.GetOrder(orderId)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve order", err.Error())
		return
	}
	// Return the order item
	utils.SuccessResponse(c, http.StatusOK, "Order retrieved successfully", order)
}

// GetAllOrders retrieves all order items
func (f *OrderController) GetAllOrders(c *gin.Context) {
	// Call the GetAllOrders service
	orders, err := f.OrderService.GetAllOrders()
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve orders", err.Error())
		return
	}

	// Return the list of orders
	utils.SuccessResponse(c, http.StatusOK, "Orders retrieved successfully", orders)
}

// GetRestaurantOrders retrieves all the orders of a particular restaurant
func (f *OrderController) GetRestaurantOrders(c *gin.Context) {
	restaurantId, valid := utils.ValidateID(c, "restaurant_id")
	if !valid {
		return
	}

	// Call the GetRestaurantOrders service
	orders, err := f.OrderService.GetRestaurantOrders(restaurantId)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve orders", err.Error())
		return
	}

	// Return the list of orders
	utils.SuccessResponse(c, http.StatusOK, "Orders retrieved successfully", orders)
}

// SearchOrder
func (f *OrderController) SearchOrder(c *gin.Context) {
	// Get the search query parameter
	query := c.Query("q")

	// Call the SearchOrder service
	orders, err := f.OrderService.SearchOrders(query)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to search for orders", err.Error())
		return
	}

	// Return the list of orders
	utils.SuccessResponse(c, http.StatusOK, "Orders retrieved successfully", orders)
}

// GetUserOrders retrieves all the orders of a particular user
func (f *OrderController) GetUserOrders(c *gin.Context) {
	userId, valid := utils.ValidateID(c, "user_id")
	if !valid {
		return
	}

	// Call the GetUserOrders service
	orders, err := f.OrderService.GetUserOrders(userId)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve orders", err.Error())
		return
	}

	// Return the list of orders
	utils.SuccessResponse(c, http.StatusOK, "Orders retrieved successfully", orders)
}

// GetOrdersByStatus retrieves all the orders of a particular status
func (f *OrderController) GetOrdersByStatus(c *gin.Context) {
	status := c.Query("status")

	// Call the GetOrdersByStatus service
	orders, err := f.OrderService.GetOrdersByStatus(status)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve orders", err.Error())
		return
	}

	// Return the list of orders
	utils.SuccessResponse(c, http.StatusOK, "Orders retrieved successfully", orders)
}
