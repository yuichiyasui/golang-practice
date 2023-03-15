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

type ErrorResponse struct {
	Message string `json:"message"`
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	user := models.User{}
	err := ctx.BindJSON(&user)
	if user.Name == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{Message: "名前は必須です"})
		return
	}
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	h.Db.Create(&user)
	ctx.JSON(http.StatusOK, &user)
}
