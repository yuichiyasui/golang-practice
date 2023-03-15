package handlers

import (
	"app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	Db *gorm.DB
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	user := models.User{}
	ctx.BindJSON(&user)
	h.Db.Create(&user)

	ctx.JSON(http.StatusOK, &user)
}
