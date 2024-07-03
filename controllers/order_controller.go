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
		// Foods        []uint  `json: "foods"`
		// Addons       []uint  `json: "addons"`
		TotalPrice float64 `json:"total_price"`
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
	// order.Foods = body.Foods
	// order.Addons = body.Addons
	order.TotalPrice = body.TotalPrice

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
		// Foods        []uint  `json: "foods"`
		// Addons       []uint  `json: "addons"`
		TotalPrice float64 `json:"total_price"`
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
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Order not found", err.Error())
		return
	}

	var order models.Order
	order.UserID = body.UserID
	order.TableID = &body.TableID
	order.RestaurantID = body.RestaurantID
	// order.Foods = body.Foods
	// order.Addons = body.Addons
	order.TotalPrice = body.TotalPrice
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
