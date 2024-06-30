package routes

import (
	"madang_api/controllers"
	"madang_api/middleware"
	"madang_api/services"

	"github.com/gin-gonic/gin"
)

func SetupRestaurantRoutes(router *gin.Engine, restaurantService *services.RestaurantService) {
	restaurantController := &controllers.RestaurantController{
		RestaurantService: services.RestaurantService{},
	}

	restaurantRoutes := router.Group("/api/restaurant")
	{
		restaurantRoutes.POST("/", middleware.AuthMiddleware, restaurantController.CreateRestaurant)
		restaurantRoutes.PUT("/:id", middleware.AuthMiddleware, restaurantController.UpdateRestaurant)
		restaurantRoutes.DELETE("/:id", middleware.AuthMiddleware, restaurantController.DeleteRestaurant)
		restaurantRoutes.GET("/", middleware.AuthMiddleware, restaurantController.GetAllRestaurant)
		restaurantRoutes.GET("/search", middleware.AuthMiddleware, restaurantController.SearchRestaurant)
		restaurantRoutes.GET("/verified", middleware.AuthMiddleware, restaurantController.GetAllVerifiedRestaurants)
		restaurantRoutes.GET("/filtered", middleware.AuthMiddleware, restaurantController.FilterRestaurant)
		restaurantRoutes.GET("/user/:user_id", middleware.AuthMiddleware, restaurantController.GetUserRestaurants)
		restaurantRoutes.GET("/:id", middleware.AuthMiddleware, restaurantController.GetRestaurant)
	}
}
