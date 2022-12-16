package routers

import (
	"todo/controllers"
	"todo/database"
	"todo/middlewares"

	"github.com/gin-gonic/gin"
)

func TaskRouter(router *gin.Engine) {
	task := router.Group("/task")
	task.Use(middlewares.AuthMiddleware(database.DbConnection))

	{
		task.GET("/", controllers.GetAllTasks)
		task.GET("/:id", controllers.GetTask)
		task.POST("/", controllers.CreateTask)
		task.PUT("/:id", controllers.UpdateTask)
		task.DELETE("/:id", controllers.DeleteTask)
	}
}