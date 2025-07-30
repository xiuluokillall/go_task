package router

import (
	"github.com/gin-gonic/gin"
	"githubgithub.com/xiuluokillall/go_task/task4/middleware"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.ErrorHandlerMiddleware())

	//apiV1 := router.Group("/api/v1")
	//{
	//
	//}
	return router
}
