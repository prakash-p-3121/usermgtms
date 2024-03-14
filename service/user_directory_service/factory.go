package user_directory_service

import (
	"github.com/prakash-p-3121/usermgtms/database"
	"github.com/prakash-p-3121/usermgtms/service/user_directory_service/impl"
)

func NewUserDirectoryService() UserDirectoryService {
	return &impl.UserDirectoryServiceImpl{DatabaseConnection: database.GetSingleStoreConnection()}
}
