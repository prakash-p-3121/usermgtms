package model

import "github.com/gin-gonic/gin"

type RestContext interface {
	/*
			A dummy function to enable GinRestContext and Future Rest API libraries
		    implement the RestContext interface as golang uses Duck-Typing
	*/
	RestCtxGet() RestContext
}

type GinRestContext struct {
	ginCtx *gin.Context
}

func (ginRestCtx *GinRestContext) CtxSet(ctx *gin.Context) {
	ginRestCtx.ginCtx = ctx
}

func (ginRestCtx *GinRestContext) CtxGet() *gin.Context {
	return ginRestCtx.ginCtx
}

func (ginRestCtx *GinRestContext) RestCtxGet() RestContext {
	return ginRestCtx
}

func NewGinRestContext(ginCtx *gin.Context) RestContext {
	ginRestCtx := new(GinRestContext)
	ginRestCtx.CtxSet(ginCtx)
	return ginRestCtx
}
