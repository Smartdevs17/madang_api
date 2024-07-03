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
	categoryService := &services.CategoryService{}
	foodService := &services.FoodService{}
	tableService := &services.TableService{}
	addonService := &services.AddonService{}
	orderService := &services.OrderService{}
	paymentService := &services.PaymentService{}
	transactionService := &services.TransactionService{}

	// Set up Gin router
	router := gin.Default()

	// Set up user routes
	routes.SetupUserRoutes(router, userService)

	//Set up restuarant routes
	routes.SetupRestaurantRoutes(router, restaurantService)

	//Set up category routes
	routes.SetupCategoryRoutes(router, categoryService)

	//Set up food routes
	routes.SetupFoodRoutes(router, foodService)

	//Set up table routes
	routes.SetupTableRoutes(router, tableService)

	//Set up addon routes
	routes.SetupAddonRoutes(router, addonService)

	//Set up order routes
	routes.SetupOrderRoutes(router, orderService)

	//Set up payment routes
	routes.SetupPaymentRoutes(router, paymentService)

	//Set up transaction routes
	routes.SetupTransactionRoutes(router, transactionService)

	fmt.Println("Server stated running successfully")

	router.Run() // listen and serve on localhost:3000
}
