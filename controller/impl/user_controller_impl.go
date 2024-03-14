package impl

import (
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/restlib"
	"github.com/prakash-p-3121/usermgtms/service/user_service"
	"github.com/prakash-p-3121/usermodel"
	"log"
)

type UserControllerImpl struct {
	UserService user_service.UserService
}

func (userControllerImpl *UserControllerImpl) UserCreate(restCtx restlib.RestContext) {

	ginRestCtx, ok := restCtx.(*restlib.GinRestContext)
	if !ok {
		internalServerErr := errorlib.NewInternalServerError("Expected GinRestContext")
		internalServerErr.SendRestResponse(ginRestCtx.CtxGet())
		return
	}

	ctx := ginRestCtx.CtxGet()
	var req usermodel.UserCreateReq
	err := ctx.BindJSON(&req)
	if err != nil {
		badReqErr := errorlib.NewBadReqError("payload-serialization")
		badReqErr.SendRestResponse(ctx)
		return
	}

	log.Println("first-name", *req.FirstName)
	log.Println("first-name", *req.LastName)
	log.Println("first-name", *req.CountryCode)
	log.Println("first-name", *req.PhoneNumberStr)
	log.Println("first-name", *req.EmailID)

	idResp, appErr := userControllerImpl.UserService.UserCreate(&req)
	if appErr != nil {
		appErr.SendRestResponse(ctx)
		return
	}

	restlib.OkNoContentResponse(ctx, idResp)
}
