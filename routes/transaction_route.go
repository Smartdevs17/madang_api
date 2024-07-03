package routes

import (
	"madang_api/controllers"
	"madang_api/middleware"
	"madang_api/services"

	"github.com/gin-gonic/gin"
)

func SetupTransactionRoutes(router *gin.Engine, transactionService *services.TransactionService) {
	transactionController := &controllers.TransactionController{
		TransactionService: services.TransactionService{},
	}

	transactionRoutes := router.Group("/api/transactions")
	{
		transactionRoutes.POST("/", middleware.AuthMiddleware, transactionController.CreateTransaction)
		transactionRoutes.PUT("/:id", middleware.AuthMiddleware, transactionController.UpdateTransaction)
		transactionRoutes.DELETE("/:id", middleware.AuthMiddleware, transactionController.DeleteTransaction)
		transactionRoutes.GET("/", middleware.AuthMiddleware, transactionController.GetAllTransactions)
		transactionRoutes.GET("/restaurant/:id", middleware.AuthMiddleware, transactionController.GetRestaurantTransactions)
		transactionRoutes.GET("/:id", middleware.AuthMiddleware, transactionController.GetTransaction)
	}
}
