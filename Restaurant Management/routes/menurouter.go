package routes

import (
	"restaurant_management/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRouters(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/menus", controllers.GetMenu())
	incomingRoutes.GET("/menus/:menus_id", controllers.GetMenu())
	incomingRoutes.POST("/menus", controllers.CreateMenu())
	incomingRoutes.PATCH("/menus:menus_id", controllers.UpdateMenu())
}