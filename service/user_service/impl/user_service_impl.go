package impl

import (
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/idgenclient"
	"github.com/prakash-p-3121/idgenmodel"
	"github.com/prakash-p-3121/usermgtms/cfg"
	"github.com/prakash-p-3121/usermgtms/database"
	"github.com/prakash-p-3121/usermgtms/repository/user_repository"
	"github.com/prakash-p-3121/usermgtms/service/user_directory_service"
	"github.com/prakash-p-3121/usermodel"
)

type UserServiceImpl struct {
	UserRepository user_repository.UserRepository
}

func (service *UserServiceImpl) CreateUser(req *usermodel.UserCreateReq) (*idgenmodel.IDGenResp, errorlib.AppError) {
	appErr := req.Validate()
	if appErr != nil {
		return nil, appErr
	}

	idGenMSCfg, err := cfg.GetMsConnectionCfg("idgenms")
	if err != nil {
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	idGenClient := idgenclient.NewIDGenClient(idGenMSCfg.Host, uint(idGenMSCfg.Port))
	resp, appErr := idGenClient.NextID(database.UsersTable)
	if appErr != nil {
		return nil, appErr
	}
	userID := resp.ID

	userDirectoryService := user_directory_service.NewUserDirectoryService()
	shardPtr, appErr := userDirectoryService.LookUpShard(userID)
	if appErr != nil {
		return nil, appErr
	}

	appErr = service.UserRepository.CreateUser(*shardPtr.ID, resp, req)
	if appErr != nil {
		return nil, appErr
	}

	return resp, nil
}

func (service *UserServiceImpl) FindUser(userID string) (*usermodel.User, errorlib.AppError) {
	userDirectoryService := user_directory_service.NewUserDirectoryService()
	shardPtr, appErr := userDirectoryService.LookUpShard(userID)
	if appErr != nil {
		return nil, appErr
	}

	user, appErr := service.UserRepository.FindUser(*shardPtr.ID, userID)
	if appErr != nil {
		return user, appErr
	}
	return user, appErr
}
