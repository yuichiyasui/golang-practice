package router

import (
	"app/handlers"

	"github.com/gin-gonic/gin"
)

func Load() *gin.Engine {
	router := gin.Default()

	todoApi := router.Group("/api/v1/todos")
	todoApi.POST("/", handlers.TodoPostHandler)

	userApi := router.Group("/api/v1/user")
	userApi.POST("/", handlers.UserPostHandler)

	return router
}
