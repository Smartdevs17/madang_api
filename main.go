package main

import (
	"fmt"
	"madang_api/config"
	"madang_api/routes"
	"madang_api/services"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVars()
	config.ConnectToDB()
	config.SyncDatabase()
}
func main() {
	userService := &services.UserService{}
	restaurantService := &services.RestaurantService{}

	// Set up Gin router
	router := gin.Default()

	// Set up user routes
	routes.SetupUserRoutes(router, userService)

	//Set up restuarant routes
	routes.SetupRestaurantRoutes(router, restaurantService)

	fmt.Println("Server stated running successfully")

	router.Run() // listen and serve on localhost:3000
}
