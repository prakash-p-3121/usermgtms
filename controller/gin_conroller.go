package controller

import (
	"github.com/gin-gonic/gin"
	"usermgtms/controller/model"
)

func UserCreate(ctx *gin.Context) {
	ginRestCtx := model.NewGinRestContext(ctx)
}
