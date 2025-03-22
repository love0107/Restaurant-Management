package routes

import (
	"restaurant_management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRouters(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/orders", controllers.GetOrders())
	incomingRoutes.GET("/orders/:order_id", controllers.GetOrder())
	incomingRoutes.POST("/orders", controllers.CreateOrders())
	incomingRoutes.PATCH("/orders:order_id", controllers.UpdateOrders())
}