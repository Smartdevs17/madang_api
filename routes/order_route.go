package routes

import (
	"madang_api/controllers"
	"madang_api/middleware"
	"madang_api/services"

	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(router *gin.Engine, orderService *services.OrderService) {
	orderController := &controllers.OrderController{
		OrderService: services.OrderService{},
	}

	orderRoutes := router.Group("/api/orders")
	{
		orderRoutes.POST("/", middleware.AuthMiddleware, orderController.AddOrder)
		orderRoutes.PUT("/:id", middleware.AuthMiddleware, orderController.UpdateOrder)
		orderRoutes.DELETE("/:id", middleware.AuthMiddleware, orderController.DeleteOrder)
		orderRoutes.GET("/", middleware.AuthMiddleware, orderController.GetAllOrders)
		orderRoutes.GET("/search", middleware.AuthMiddleware, orderController.SearchOrder)
		orderRoutes.GET("/status", middleware.AuthMiddleware, orderController.GetOrdersByStatus)
		orderRoutes.GET("/restaurant/:restaurant_id", middleware.AuthMiddleware, orderController.GetRestaurantOrders)
		orderRoutes.GET("/user/:user_id", middleware.AuthMiddleware, orderController.GetUserOrders)
		orderRoutes.GET("/:id", middleware.AuthMiddleware, orderController.GetOrder)
	}
}
