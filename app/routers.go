package main

import (
	"github.com/gin-gonic/gin"
)

func Load() *gin.Engine {
	router := gin.Default()

	todoApi := router.Group("/api/v1/todos")
	todoApi.POST("/", TodoPostHandler)

	return router
}
