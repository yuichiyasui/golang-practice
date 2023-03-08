package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserPostHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Create user!",
	})
}
