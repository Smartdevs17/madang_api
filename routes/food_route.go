package routes

import (
	"madang_api/controllers"
	"madang_api/middleware"
	"madang_api/services"

	"github.com/gin-gonic/gin"
)

func SetupFoodRoutes(router *gin.Engine, foodService *services.FoodService) {
	foodController := &controllers.FoodController{
		FoodService: services.FoodService{},
	}

	foodRoutes := router.Group("/api/foods")
	{
		foodRoutes.POST("/", middleware.AuthMiddleware, foodController.AddFood)
		foodRoutes.PUT("/:id", middleware.AuthMiddleware, foodController.UpdateFood)
		foodRoutes.DELETE("/:id", middleware.AuthMiddleware, foodController.DeleteFood)
		foodRoutes.GET("/", middleware.AuthMiddleware, foodController.GetAllFoods)
		foodRoutes.GET("/search", middleware.AuthMiddleware, foodController.SearchFood)
		foodRoutes.GET("/restaurant/:id", middleware.AuthMiddleware, foodController.GetRestaurantFoods)
		foodRoutes.GET("/:id", middleware.AuthMiddleware, foodController.GetFood)
	}
}
