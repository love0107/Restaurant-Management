package routes

import (
	"restaurant_management/controllers"

	"github.com/gin-gonic/gin"
)


func OrderRouters(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/orderItems", controllers.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id", controllers.GetOrderItem())
	incomingRoutes.GET("/orderItem-order/order_id", controllers.OrderItemsByOrder())
	incomingRoutes.POST("/orderItems", controllers.CreateOrdersItem())
	incomingRoutes.PATCH("/orderItems:orderItem_id", controllers.UpdateOrderItem())
}