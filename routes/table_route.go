package routes

import (
	"madang_api/controllers"
	"madang_api/middleware"
	"madang_api/services"

	"github.com/gin-gonic/gin"
)

func SetupTableRoutes(router *gin.Engine, tableService *services.TableService) {
	tableController := &controllers.TableController{
		TableService: services.TableService{},
	}

	tableRoutes := router.Group("/api/tables")
	{
		tableRoutes.POST("/", middleware.AuthMiddleware, tableController.AddTable)
		tableRoutes.PUT("/:id", middleware.AuthMiddleware, tableController.UpdateTable)
		tableRoutes.DELETE("/:id", middleware.AuthMiddleware, tableController.DeleteTable)
		tableRoutes.GET("/", middleware.AuthMiddleware, tableController.GetAllTables)
		tableRoutes.GET("/search", middleware.AuthMiddleware, tableController.SearchTable)
		tableRoutes.GET("/restaurant/:id", middleware.AuthMiddleware, tableController.GetRestaurantTables)
		tableRoutes.GET("/:id", middleware.AuthMiddleware, tableController.GetTable)
	}
}
