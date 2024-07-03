package routes

import (
	"madang_api/controllers"
	"madang_api/middleware"
	"madang_api/services"

	"github.com/gin-gonic/gin"
)

func SetupCategoryRoutes(router *gin.Engine, categoryService *services.CategoryService) {
	categoryController := &controllers.CategoryController{
		CategoryService: services.CategoryService{},
	}

	categoryRoutes := router.Group("/api/categories")
	{
		categoryRoutes.POST("/", middleware.AuthMiddleware, categoryController.CreateCategory)
		categoryRoutes.PUT("/:id", middleware.AuthMiddleware, categoryController.UpdateCategory)
		categoryRoutes.DELETE("/:id", middleware.AuthMiddleware, categoryController.DeleteCategory)
		categoryRoutes.GET("/", middleware.AuthMiddleware, categoryController.GetAllCategories)
		categoryRoutes.GET("/restaurant/:id", middleware.AuthMiddleware, categoryController.GetRestaurantCategories)
		categoryRoutes.GET("/:id", middleware.AuthMiddleware, categoryController.GetCategory)
	}
}
