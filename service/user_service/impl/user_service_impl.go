package impl

import (
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/idgenclient"
	restlib_model "github.com/prakash-p-3121/restlib/model"
	"github.com/prakash-p-3121/usermgtms/cfg"
	"github.com/prakash-p-3121/usermgtms/database"
	"github.com/prakash-p-3121/usermgtms/repository/user_repository"
	"github.com/prakash-p-3121/usermgtms/service/user_directory_service"
	"github.com/prakash-p-3121/usermodel"
)

type UserServiceImpl struct {
	UserRepository user_repository.UserRepository
}

func (service *UserServiceImpl) UserCreate(req *usermodel.UserCreateReq) (*restlib_model.IDResponse, errorlib.AppError) {
	appErr := req.Validate()
	if appErr != nil {
		return nil, appErr
	}

	idGenMSCfg, err := cfg.GetMsConnectionCfg("idgenms")
	if err != nil {
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	idGenClient := idgenclient.NewIDGenClient(idGenMSCfg.Host, idGenMSCfg.Port)
	resp, appErr := idGenClient.NextID(database.UsersTable)
	if appErr != nil {
		return nil, appErr
	}
	userID := resp.ID

	userDirectoryService := user_directory_service.NewUserDirectoryService()
	writeShardPtr, appErr := userDirectoryService.LookUpCurrentWriteShard(userID)
	if appErr != nil {
		return nil, appErr
	}

	appErr = service.UserRepository.UserCreate(*writeShardPtr.ID, resp, req)
	if appErr != nil {
		return nil, appErr
	}

	return &restlib_model.IDResponse{ResourceID: userID}, nil
}
