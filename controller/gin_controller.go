package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prakash-p-3121/restlib"
)

func CreateUser(ctx *gin.Context) {
	ginRestCtx := restlib.NewGinRestContext(ctx)
	controller := NewUserController()
	controller.CreateUser(ginRestCtx)
}

func FindUser(ctx *gin.Context) {
	ginRestCtx := restlib.NewGinRestContext(ctx)
	controller := NewUserController()
	controller.FindUser(ginRestCtx)
}
