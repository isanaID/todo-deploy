package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	UserRouter(router)
	TaskRouter(router)
	StatusTaskRouter(router)
	CategoryRouter(router)

	return router
}