package routers

import (
	"todo/controllers"
	"todo/database"
	"todo/middlewares"

	"github.com/gin-gonic/gin"
)

func StatusTaskRouter(router *gin.Engine) {
	statusTask := router.Group("/status-task")
	statusTask.Use(middlewares.AuthMiddleware(database.DbConnection))
	{
		statusTask.GET("/", controllers.GetAllStatusTasks)
		statusTask.GET("/:id", controllers.GetStatusTask)
		statusTask.POST("/", controllers.CreateStatusTask)
		statusTask.PUT("/:id", controllers.UpdateStatusTask)
		statusTask.DELETE("/:id", controllers.DeleteStatusTask)
	}
}