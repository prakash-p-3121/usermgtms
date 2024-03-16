package impl

import (
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/restlib"
	"github.com/prakash-p-3121/usermgtms/service/user_service"
	"github.com/prakash-p-3121/usermodel"
	"strings"
)

type UserControllerImpl struct {
	UserService user_service.UserService
}

func (userControllerImpl *UserControllerImpl) CreateUser(restCtx restlib.RestContext) {

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

	idResp, appErr := userControllerImpl.UserService.CreateUser(&req)
	if appErr != nil {
		appErr.SendRestResponse(ctx)
		return
	}

	restlib.OkResponse(ctx, idResp)
}

func (userControllerImpl *UserControllerImpl) FindUser(restCtx restlib.RestContext) {

	ginRestCtx, ok := restCtx.(*restlib.GinRestContext)
	if !ok {
		internalServerErr := errorlib.NewInternalServerError("Expected GinRestContext")
		internalServerErr.SendRestResponse(ginRestCtx.CtxGet())
		return
	}

	ctx := ginRestCtx.CtxGet()
	userIDStr := ctx.Query("id")
	userIDStr = strings.TrimSpace(userIDStr)
	if len(userIDStr) == 0 {
		badReqErr := errorlib.NewBadReqError("id")
		badReqErr.SendRestResponse(ginRestCtx.CtxGet())
		return
	}

	idResp, appErr := userControllerImpl.UserService.FindUser(userIDStr)
	if appErr != nil {
		appErr.SendRestResponse(ctx)
		return
	}

	restlib.OkResponse(ctx, idResp)
}
