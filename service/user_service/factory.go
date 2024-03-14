package user_service

import (
	"database/sql"
	"github.com/prakash-p-3121/usermgtms/repository/user_repository"
	"github.com/prakash-p-3121/usermgtms/service/user_service/impl"
	"sync"
)

func NewUserService(singleStoreConnection *sql.DB, shardConnectionsMap *sync.Map) UserService {
	repository := user_repository.NewUserRepository(singleStoreConnection, shardConnectionsMap)
	return &impl.UserServiceImpl{UserRepository: repository}
}
