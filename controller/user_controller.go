package controller

import "github.com/prakash-p-3121/usermgtms/controller/model"

type UserController interface {
	UserCreate(restCtx model.RestContext)
	//UserGet(ctx model.RestContext) error
	//UserUpdate(ctx model.RestContext) error
	//UserDelete(ctx model.RestContext) error
}
