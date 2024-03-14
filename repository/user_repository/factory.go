package user_repository

import (
	"database/sql"
	"github.com/prakash-p-3121/usermgtms/repository/user_repository/impl"
	"sync"
)

func NewUserRepository(singleStoreConnection *sql.DB, shardConnectionsMap *sync.Map) UserRepository {
	return &impl.UserRepositoryImpl{ShardConnectionsMap: shardConnectionsMap, SingleStoreConnection: singleStoreConnection}
}
