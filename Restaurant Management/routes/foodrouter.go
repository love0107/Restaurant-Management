package routes

import (
	"restaurant_management/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRouters(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/foods", controllers.GetFoods())
	incomingRoutes.GET("/foods/:food_id", controllers.GetFood())
	incomingRoutes.POST("/foods", controllers.CreateFood())
	incomingRoutes.PATCH("/foods:food_id", controllers.UpdateFood())
}