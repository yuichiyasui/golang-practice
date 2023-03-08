package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TodoPostHandler(ctx *gin.Context) {
	println("Create task!")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Create task!",
	})
}
