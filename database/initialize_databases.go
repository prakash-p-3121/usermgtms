package database

import (
	"database/sql"
	"sync"
)

var shardIDToDatabaseConnectionMap *sync.Map
var singleStoreDatabaseConnection *sql.DB

const (
	UsersTable string = "users"
)

func SetShardConnectionsMap(connectionsMap *sync.Map) {
	shardIDToDatabaseConnectionMap = connectionsMap
}

func GetShardConnectionsMap() *sync.Map {
	return shardIDToDatabaseConnectionMap
}

func SetSingleStoreConnection(databaseConnection *sql.DB) {
	singleStoreDatabaseConnection = databaseConnection
}

func GetSingleStoreConnection() *sql.DB {
	return singleStoreDatabaseConnection
}

func GetShardedTableList() []string {
	return []string{UsersTable}
}
