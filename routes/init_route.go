package routes

import (
	"madang_api/controllers"
	"madang_api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupInitRoutes(router *gin.Engine) {
	initController := &controllers.InitController{}

	initRoutes := router.Group("/api/inits")
	{
		initRoutes.GET("/", middleware.AuthMiddleware, initController.LoadData)
	}
}
