package routes

import (
	"restaurant_management/controllers"

	"github.com/gin-gonic/gin"
)

func TableRouters(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/tables", controllers.GetTables())
	incomingRoutes.GET("/tables/:table_id", controllers.GetTable())
	incomingRoutes.POST("/tables", controllers.CreateTables())
	incomingRoutes.PATCH("/tables:table_id", controllers.UpdateTables())
}