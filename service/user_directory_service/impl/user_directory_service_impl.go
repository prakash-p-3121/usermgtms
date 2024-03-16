package impl

import (
	"database/sql"
	database_clustermgt_client "github.com/prakash-p-3121/database-clustermgt-client"
	model "github.com/prakash-p-3121/database-clustermgt-model"
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/usermgtms/cfg"
	"github.com/prakash-p-3121/usermgtms/database"
	"github.com/prakash-p-3121/usermgtms/repository/user_directory_repository"
	"log"
)

type UserDirectoryServiceImpl struct {
	DatabaseConnection      *sql.DB
	UserDirectoryRepository user_directory_repository.UserDirectoryRepository
}

func (userService *UserDirectoryServiceImpl) LookUpByEmailID(emailID string) (*model.EmailIDLookUpResp,
	errorlib.AppError) {
	repo := userService.UserDirectoryRepository
	return repo.LookUpByEmailID(emailID)
}

func (userService *UserDirectoryServiceImpl) LookUpByUserID(userID string) (*model.UserIDLookUpResp, errorlib.AppError) {
	repo := userService.UserDirectoryRepository
	return repo.LookUpByUserID(userID)
}

func (userService *UserDirectoryServiceImpl) LookUpShard(userID string) (*model.DatabaseShard, errorlib.AppError) {
	databaseClstrMgtMsCfg, err := cfg.GetMsConnectionCfg("database-clustermgt-ms")
	if err != nil {
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	client := database_clustermgt_client.NewDatabaseClusterMgtClient(databaseClstrMgtMsCfg.Host, uint(databaseClstrMgtMsCfg.Port))
	shardPtr, appErr := client.FindShard(database.UsersTable, userID)
	if appErr != nil {
		log.Println("error-received")
		return nil, appErr
	} else {
		log.Println("no-errors")
	}
	log.Println("shardPtr")
	log.Println(*shardPtr)
	return shardPtr, nil
}
