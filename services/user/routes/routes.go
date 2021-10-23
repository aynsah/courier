package routes

import (
	"courier/services/user/controllers"

	"github.com/gin-gonic/gin"
)

func V1Application(router *gin.Engine) {
	v1 := router.Group("api/v1")
	{
		users := v1.Group("users")
		{
			users.POST("/sign-up", controllers.SignUp)
			users.POST("/login", controllers.Login)
			users.POST("/info", controllers.Info)
		}
	}
}
