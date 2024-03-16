package user_service

import (
	"github.com/prakash-p-3121/errorlib"
	restlib_model "github.com/prakash-p-3121/restlib/model"
	"github.com/prakash-p-3121/usermodel"
)

type UserService interface {
	CreateUser(req *usermodel.UserCreateReq) (*restlib_model.IDResponse, errorlib.AppError)
	FindUser(userID string) (*usermodel.User, errorlib.AppError)
}
