package handler

import (
	"github.com/gin-gonic/gin"
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/infrastructure/middleware"
	"volunteer-team/backend/internal/infrastructure/types"
	"volunteer-team/backend/internal/infrastructure/utils/response"
	"volunteer-team/backend/internal/logic"
)

type UserHandler struct {
	userLogic *logic.UserLogic
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		userLogic: logic.NewUserLogic(),
	}
}

func (uh *UserHandler) Login(c *gin.Context) {
	cr := middleware.GetBind[types.LoginReq](c)
	global.Log.Info(cr)
	resp, err := uh.userLogic.LoginLogic(cr)
	response.Response(c, resp, err)
}

func (uh *UserHandler) Register(c *gin.Context) {
	cr := middleware.GetBind[types.RegisterReq](c)
	global.Log.Info(cr)
	resp, err := uh.userLogic.RegisterLogic(cr)
	response.Response(c, resp, err)
}

func (uh *UserHandler) GetLoginCode(c *gin.Context) {
	cr := middleware.GetBind[types.GetCodeReq](c)
	global.Log.Info(cr)
	resp, err := uh.userLogic.GetLoginCode(cr)
	response.Response(c, resp, err)
}

func (uh *UserHandler) GetRegisterCode(c *gin.Context) {
	cr := middleware.GetBind[types.GetCodeReq](c)
	global.Log.Info(cr)
	resp, err := uh.userLogic.GetRegisterCode(cr)
	response.Response(c, resp, err)
}

func (uh *UserHandler) GetResetCode(c *gin.Context) {
	cr := middleware.GetBind[types.GetCodeReq](c)
	global.Log.Info(cr)
	resp, err := uh.userLogic.GetResetCode(cr)
	response.Response(c, resp, err)
}
