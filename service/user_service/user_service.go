package user_service

import (
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/idgenmodel"
	"github.com/prakash-p-3121/usermodel"
)

type UserService interface {
	CreateUser(req *usermodel.UserCreateReq) (*idgenmodel.IDGenResp, errorlib.AppError)
	FindUser(userID string) (*usermodel.User, errorlib.AppError)
}
