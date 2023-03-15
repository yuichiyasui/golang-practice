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
	taskHandler := handlers.TaskHandler{
		Db: dbConnection,
	}

	router := gin.Default()

	taskApi := router.Group("/api/v1/task")
	taskApi.POST("/", taskHandler.CreateTask)

	userApi := router.Group("/api/v1/user")
	userApi.POST("/", userHandler.CreateUser)

	return router
}
