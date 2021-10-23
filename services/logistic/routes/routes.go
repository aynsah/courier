package routes

import (
	"courier/services/logistic/controllers"
	"courier/services/logistic/middleware"

	"github.com/gin-gonic/gin"
)

func V1Application(router *gin.Engine) {
	v1 := router.Group("api/v1")
	{
		logistics := v1.Group("logistics")
		{
			logistics.Use(middleware.Verify)
			logistics.POST("/", controllers.LogisticList)
			logistics.POST("/search", controllers.SearchLogistics)
		}
	}
}
