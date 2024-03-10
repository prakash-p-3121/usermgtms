package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prakash-p-3121/usermgtms/controller/model"
)

func UserCreate(ctx *gin.Context) {
	ginRestCtx := model.NewGinRestContext(ctx)
	controller := NewUserController()
	controller.UserCreate(ginRestCtx)
}
