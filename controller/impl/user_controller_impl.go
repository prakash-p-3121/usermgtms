package impl

import (
	"github.com/prakash-p-3121/errorlib"
	rest_response_lib "github.com/prakash-p-3121/rest-response-lib"
	"github.com/prakash-p-3121/usermgtms/controller/model"
	"github.com/prakash-p-3121/usermodel"
	"log"
)

type UserControllerImpl struct {
}

func (userControllerImpl *UserControllerImpl) UserCreate(restCtx model.RestContext) {

	ginRestCtx, ok := restCtx.(*model.GinRestContext)
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

	rest_response_lib.OkNoContentResponse(ctx, nil)
}
