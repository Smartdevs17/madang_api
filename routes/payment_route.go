package routes

import (
	"madang_api/controllers"
	"madang_api/middleware"
	"madang_api/services"

	"github.com/gin-gonic/gin"
)

func SetupPaymentRoutes(router *gin.Engine, paymentService *services.PaymentService) {
	paymentController := &controllers.PaymentController{
		PaymentService: services.PaymentService{},
	}

	paymentRoutes := router.Group("/api/payments")
	{
		paymentRoutes.POST("/", middleware.AuthMiddleware, paymentController.CreatePayment)
		paymentRoutes.PUT("/:id", middleware.AuthMiddleware, paymentController.UpdatePayment)
		paymentRoutes.DELETE("/:id", middleware.AuthMiddleware, paymentController.DeletePayment)
		paymentRoutes.GET("/", middleware.AuthMiddleware, paymentController.GetAllPayments)
		paymentRoutes.GET("/restaurant/:id", middleware.AuthMiddleware, paymentController.GetRestaurantPayments)
		paymentRoutes.GET("/:id", middleware.AuthMiddleware, paymentController.GetPayment)
	}
}
