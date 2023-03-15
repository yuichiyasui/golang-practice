package handlers

import (
	"app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TaskHandler struct {
	Db *gorm.DB
}

func (h *TaskHandler) CreateTask(ctx *gin.Context) {
	task := models.Task{}
	err := ctx.BindJSON(&task)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	h.Db.Create(&task)
	ctx.JSON(http.StatusOK, &task)
}
