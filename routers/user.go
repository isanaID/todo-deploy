package routers

import (
	"todo/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	user := router.Group("/user")
	{
		user.POST("/register", controllers.RegisterUser)
		user.POST("/login", controllers.LoginUser)
	}
}