package controllers

import (
	"madang_api/models"
	"madang_api/services"
	"madang_api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	PaymentService services.PaymentService
}

type PaymentControllerInterface interface {
	GetAllPayments(c *gin.Context)
	GetPayment(c *gin.Context)
	CreatePayment(c *gin.Context)
	UpdatePayment(c *gin.Context)
	DeletePayment(c *gin.Context)
	GetRestaurantPayments(c *gin.Context)
}

func (ctrl *PaymentController) GetAllPayments(c *gin.Context) {
	payments, err := ctrl.PaymentService.GetPayments()
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve payments", err.Error())
		return
	}

	// Return the list of payments
	utils.SuccessResponse(c, http.StatusOK, "Payments retrieved successfully", payments)
}

func (ctrl *PaymentController) GetPayment(c *gin.Context) {
	// Get the payment ID from the request parameters
	paymentID, valid := utils.ValidateID(c, "id")
	if !valid {
		return
	}

	// Call the service to retrieve the payment
	payment, err := ctrl.PaymentService.GetPayment(paymentID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve payment", err.Error())
		return
	}

	// Return the payment
	utils.SuccessResponse(c, http.StatusOK, "Payment retrieved successfully", payment)
}

func (ctrl *PaymentController) CreatePayment(c *gin.Context) {
	// Bind the request body to a Payment struct
	var body struct {
		OrderID      uint    `json:"order_id"`
		Method       string  `json:"method"`
		Status       string  `json:"status"`
		Amount       float64 `json:"amount"`
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
	var payment models.Payment
	payment.OrderID = body.OrderID
	payment.Method = body.Method
	payment.Status = body.Status
	payment.Amount = body.Amount
	payment.RestaurantID = body.RestaurantID

	// Call the service to create the payment
	newPayment, err := ctrl.PaymentService.CreatePayment(&payment)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create payment", err.Error())
		return
	}

	// Return the created payment
	utils.SuccessResponse(c, http.StatusCreated, "Payment created successfully", newPayment)
}

func (ctrl *PaymentController) UpdatePayment(c *gin.Context) {
	// Get the payment ID from the request parameters
	paymentID, valid := utils.ValidateID(c, "id")
	if !valid {
		return
	}

	// Bind the request body to a Payment struct
	var body struct {
		OrderID      uint    `json:"order_id"`
		Method       string  `json:"method"`
		Status       string  `json:"status"`
		Amount       float64 `json:"amount"`
		RestaurantID uint    `json:"restaurant_id"`
	}

	// Validate the request body
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}
	//check fig the payment exists
	_, err := ctrl.PaymentService.GetPayment(paymentID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Payment not found", err.Error())
		return
	}

	var payment models.Payment
	payment.OrderID = body.OrderID
	payment.Method = body.Method
	payment.Status = body.Status
	payment.Amount = body.Amount
	payment.RestaurantID = body.RestaurantID

	// Call the service to update the payment
	updatedPayment, err := ctrl.PaymentService.UpdatePayment(&payment)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update payment", err.Error())
		return
	}

	// Return the updated payment
	utils.SuccessResponse(c, http.StatusOK, "Payment updated successfully", updatedPayment)
}

func (ctrl *PaymentController) DeletePayment(c *gin.Context) {
	// Get the payment ID from the request parameters
	paymentID, valid := utils.ValidateID(c, "id")
	if !valid {
		return
	}

	//check if the payment exists
	_, errExist := ctrl.PaymentService.GetPayment(paymentID)
	// Handle error
	if errExist != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Payment not found", errExist.Error())
		return
	}

	// Call the service to delete the payment
	err := ctrl.PaymentService.DeletePayment(paymentID)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete payment", err.Error())
		return
	}

	// Return a success response
	utils.SuccessResponse(c, http.StatusOK, "Payment deleted successfully", nil)
}

func (ctrl *PaymentController) GetRestaurantPayments(c *gin.Context) {
	// Get the restaurant ID from the request parameters
	restaurantID, valid := utils.ValidateID(c, "restaurant_id")
	if !valid {
		return
	}

	// Call the service to retrieve the payments for the restaurant
	payments, err := ctrl.PaymentService.GetRestaurantPayments(restaurantID)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve payments", err.Error())
		return
	}

	// Return the list of payments
	utils.SuccessResponse(c, http.StatusOK, "Payments retrieved successfully", payments)
}
