package routes

import (
	"madang_api/controllers"
	"madang_api/middleware"
	"madang_api/services"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine, userService *services.UserService) {
	userController := controllers.UserController{UserService: userService}

	router.POST("/api/auth/register", userController.Register)
	router.POST("/api/auth/verify-email", userController.ValidateEmail)
	router.POST("/api/auth/login", userController.Login)
	router.GET("/api/users/:id", middleware.AuthMiddleware, userController.GetUserByID)
	router.GET("/api/users", middleware.AuthMiddleware, userController.GetAllUsers)
	router.PUT("/api/users/:id", middleware.AuthMiddleware, userController.UpdateUser)
	router.DELETE("/api/users/:id", middleware.AuthMiddleware, userController.DeleteUser)

}
