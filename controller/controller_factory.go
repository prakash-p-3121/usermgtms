package controller

import (
	"github.com/prakash-p-3121/usermgtms/controller/impl"
	"github.com/prakash-p-3121/usermgtms/database"
	"github.com/prakash-p-3121/usermgtms/service/user_service"
)

func NewUserController() UserController {
	service := user_service.NewUserService(database.GetSingleStoreConnection(), database.GetShardConnectionsMap())
	return &impl.UserControllerImpl{UserService: service}
}
