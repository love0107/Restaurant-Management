package routes

import (
	"restaurant_management/controllers"

	"github.com/gin-gonic/gin"
)

func InvoceRouters(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/invoices", controllers.GetInvoices())
	incomingRoutes.GET("/invoices/:invoice_id", controllers.GetInvoice())
	incomingRoutes.POST("/invoices", controllers.CreateFood())
	incomingRoutes.PATCH("/invoices:invoice_id", controllers.UpdateInvoice())
}