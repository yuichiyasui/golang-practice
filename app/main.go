package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/task", func(ctx *gin.Context) {
		println("Create task!")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Create task!",
		})
	})

	router.Run()
}
