package controller

import "github.com/prakash-p-3121/usermgtms/controller/impl"

func NewUserController() UserController {
	return &impl.UserControllerImpl{}
}
