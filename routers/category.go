package routers

import (
	"todo/controllers"
	"todo/database"
	"todo/middlewares"

	"github.com/gin-gonic/gin"
)

func CategoryRouter(router *gin.Engine) {
	category := router.Group("/category")
	category.Use(middlewares.AuthMiddleware(database.DbConnection))
	{
		category.GET("/", controllers.GetAllCategories)
		category.GET("/:id", controllers.GetCategory)
		category.POST("/", controllers.CreateCategory)
		category.PUT("/:id", controllers.UpdateCategory)
		category.DELETE("/:id", controllers.DeleteCategory)
	}
}