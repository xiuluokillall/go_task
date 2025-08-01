package router

import (
	"github.com/gin-gonic/gin"
	"githubgithub.com/xiuluokillall/go_task/task4/internal/handler"
	"githubgithub.com/xiuluokillall/go_task/task4/middleware"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.ErrorHandlerMiddleware())

	apiV1 := router.Group("/api/v1")
	{
		authGroup := apiV1.Group("auth")
		{
			authGroup.POST("/register", handler.Register)
			authGroup.POST("/login", handler.Login)
		}

		postGroup := apiV1.Group("/post").Use(middleware.JwtAuth())
		{
			postGroup.POST("/create", handler.CreatePost)
		}
	}
	return router
}
