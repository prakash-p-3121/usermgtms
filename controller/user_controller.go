package controller

import "github.com/prakash-p-3121/restlib"

type UserController interface {
	CreateUser(restCtx restlib.RestContext)
	FindUser(restCtx restlib.RestContext)
	//UserGet(ctx model.RestContext) error
	//UserUpdate(ctx model.RestContext) error
	//UserDelete(ctx model.RestContext) error
}
