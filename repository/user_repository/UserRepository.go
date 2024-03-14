package user_repository

import (
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/idgenmodel"
	"github.com/prakash-p-3121/usermodel"
)

type UserRepository interface {
	UserCreate(shardID int64, idGenResp *idgenmodel.IDGenResp, req *usermodel.UserCreateReq) errorlib.AppError
}
