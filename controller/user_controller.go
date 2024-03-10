package controller

import "usermgtms/controller/model"

type UserController interface {
	UserCreate(ctx model.RestContext) error
	UserGet(ctx model.RestContext) error
	UserUpdate(ctx model.RestContext) error
	UserDelete(ctx model.RestContext) error
}
