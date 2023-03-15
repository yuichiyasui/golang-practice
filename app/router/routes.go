package router

import (
	"app/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Load(dbConnection *gorm.DB) *gin.Engine {
	userHandler := handlers.UserHandler{
		Db: dbConnection,
	}

	router := gin.Default()

	todoApi := router.Group("/api/v1/todos")
	todoApi.POST("/", handlers.TodoPostHandler)

	userApi := router.Group("/api/v1/user")
	userApi.POST("/", userHandler.CreateUser)

	return router
}
