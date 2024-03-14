package user_directory_service

import (
	model "github.com/prakash-p-3121/database-clustermgt-model"
	"github.com/prakash-p-3121/errorlib"
)

type UserDirectoryService interface {
	LookUpByEmailID(emailID string) (*model.EmailIDLookUpResp, errorlib.AppError)
	LookUpByUserID(userID string) (*model.UserIDLookUpResp, errorlib.AppError)
	LookUpCurrentWriteShard(userID string) (*model.DatabaseShard, errorlib.AppError)
}
