package routes

import (
	"madang_api/controllers"
	"madang_api/middleware"
	"madang_api/services"

	"github.com/gin-gonic/gin"
)

func SetupAddonRoutes(router *gin.Engine, addonService *services.AddonService) {
	addonController := &controllers.AddonController{
		AddonService: services.AddonService{},
	}

	addonRoutes := router.Group("/api/addons")
	{
		addonRoutes.POST("/", middleware.AuthMiddleware, addonController.AddAddon)
		addonRoutes.PUT("/:id", middleware.AuthMiddleware, addonController.UpdateAddon)
		addonRoutes.DELETE("/:id", middleware.AuthMiddleware, addonController.DeleteAddon)
		addonRoutes.GET("/", middleware.AuthMiddleware, addonController.GetAllAddons)
		addonRoutes.GET("/search", middleware.AuthMiddleware, addonController.SearchAddon)
		addonRoutes.GET("/restaurant/:id", middleware.AuthMiddleware, addonController.GetRestaurantAddons)
		addonRoutes.GET("/:id", middleware.AuthMiddleware, addonController.GetAddon)
	}
}
