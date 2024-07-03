package controllers

import (
	"madang_api/models"
	"madang_api/services"
	"madang_api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	TransactionService services.TransactionService
}

type TransactionControllerInterface interface {
	GetAllTransactions(c *gin.Context)
	GetTransaction(c *gin.Context)
	CreateTransaction(c *gin.Context)
	UpdateTransaction(c *gin.Context)
	DeleteTransaction(c *gin.Context)
	GetRestaurantTransactions(c *gin.Context)
}

func (ctrl *TransactionController) GetAllTransactions(c *gin.Context) {
	transactions, err := ctrl.TransactionService.GetTransactions()
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve transactions", err.Error())
		return
	}

	// Return the list of transactions
	utils.SuccessResponse(c, http.StatusOK, "Transactions retrieved successfully", transactions)
}

func (ctrl *TransactionController) GetTransaction(c *gin.Context) {
	// Get the transaction ID from the request parameters
	transactionID, valid := utils.ValidateID(c, "id")
	if !valid {
		return
	}

	// Call the service to retrieve the transaction
	transaction, err := ctrl.TransactionService.GetTransaction(transactionID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve transaction", err.Error())
		return
	}

	// Return the transaction
	utils.SuccessResponse(c, http.StatusOK, "Transaction retrieved successfully", transaction)
}

func (ctrl *TransactionController) CreateTransaction(c *gin.Context) {
	// Bind the request body to a Transaction struct
	var body struct {
		OrderID      uint    `json:"order_id"`
		PaymentID    uint    `json:"payment_id"`
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
	var transaction models.Transaction
	transaction.OrderID = body.OrderID
	transaction.PaymentID = body.PaymentID
	transaction.Status = body.Status
	transaction.Amount = body.Amount
	transaction.RestaurantID = body.RestaurantID

	// Call the service to create the transaction
	newTransaction, err := ctrl.TransactionService.CreateTransaction(&transaction)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create transaction", err.Error())
		return
	}

	// Return the created transaction
	utils.SuccessResponse(c, http.StatusCreated, "Transaction created successfully", newTransaction)
}

func (ctrl *TransactionController) UpdateTransaction(c *gin.Context) {
	// Get the transaction ID from the request parameters
	transactionID, valid := utils.ValidateID(c, "id")
	if !valid {
		return
	}

	// Bind the request body to a Transaction struct
	var body struct {
		OrderID      uint    `json:"order_id"`
		PaymentID    uint    `json:"payment_id"`
		Status       string  `json:"status"`
		Amount       float64 `json:"amount"`
		RestaurantID uint    `json:"restaurant_id"`
	}

	// Validate the request body
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}
	//check fig the transaction exists
	_, err := ctrl.TransactionService.GetTransaction(transactionID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Transaction not found", err.Error())
		return
	}
	var transaction models.Transaction
	transaction.OrderID = body.OrderID
	transaction.PaymentID = body.PaymentID
	transaction.Status = body.Status
	transaction.Amount = body.Amount
	transaction.RestaurantID = body.RestaurantID

	// Call the service to update the transaction
	updatedTransaction, err := ctrl.TransactionService.UpdateTransaction(&transaction)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update transaction", err.Error())
		return
	}

	// Return the updated transaction
	utils.SuccessResponse(c, http.StatusOK, "Transaction updated successfully", updatedTransaction)
}

func (ctrl *TransactionController) DeleteTransaction(c *gin.Context) {
	// Get the transaction ID from the request parameters
	transactionID, valid := utils.ValidateID(c, "id")
	if !valid {
		return
	}

	//check if the transaction exists
	_, errExist := ctrl.TransactionService.GetTransaction(transactionID)
	// Handle error
	if errExist != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Transaction not found", errExist.Error())
		return
	}

	// Call the service to delete the transaction
	err := ctrl.TransactionService.DeleteTransaction(transactionID)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete transaction", err.Error())
		return
	}

	// Return a success response
	utils.SuccessResponse(c, http.StatusOK, "Transaction deleted successfully", nil)
}

func (ctrl *TransactionController) GetRestaurantTransactions(c *gin.Context) {
	// Get the restaurant ID from the request parameters
	restaurantID, valid := utils.ValidateID(c, "restaurant_id")
	if !valid {
		return
	}

	// Call the service to retrieve the transactions for the restaurant
	transactions, err := ctrl.TransactionService.GetRestaurantTransactions(restaurantID)
	// Handle error
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve transactions", err.Error())
		return
	}

	// Return the list of transactions
	utils.SuccessResponse(c, http.StatusOK, "Transactions retrieved successfully", transactions)
}
