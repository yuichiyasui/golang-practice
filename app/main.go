package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/task", TodoPostHandler)

	router.Run()
}
