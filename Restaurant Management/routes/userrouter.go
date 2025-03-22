package routes

import (
	"restaurant_management/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouters(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
	incomingRoutes.POST("/users/signup", controllers.singUp())
	incomingRoutes.POST("/users/login", controllers.LongIn())
}