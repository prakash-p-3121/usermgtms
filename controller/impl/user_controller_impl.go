package impl

import (
	"errors"
	"usermgtms/controller/model"
)

type UserControllerImpl struct {
}

func (userControllerImpl *UserControllerImpl) UserContext(restCtx model.RestContext) error {
	ginRestCtx, ok := restCtx.(*model.GinRestContext)
	if !ok {
		return errors.New("Expected GinRestContext")
	}
	ctx := ginRestCtx.CtxGet()
	ctx.BindJSON()
}
