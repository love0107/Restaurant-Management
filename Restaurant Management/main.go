package main

import (
	"os"
	"restaurant_management/database"
	"restaurant_management/middleware"
	"restaurant_management/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)
var foodCollection *mongo.Collection =database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(middleware.Authenticantion())

	routes.UserRouters(router)
	routes.FoodRouters(router)
	routes.MenuRouters(router)
	routes.TableRouters(router)
	routes.OrderRouters(router)
	routes.OrderItemRouters(router)
	routes.InvoceRouters(router)
}
