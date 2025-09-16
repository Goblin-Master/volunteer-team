package handler

import (
	"fmt"
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/infrastructure/middleware"
	"volunteer-team/backend/internal/infrastructure/pkg/jwtx"
	"volunteer-team/backend/internal/infrastructure/types"
	"volunteer-team/backend/internal/infrastructure/utils/response"
	"volunteer-team/backend/internal/logic"

	"github.com/gin-gonic/gin"
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
	resp, err := uh.userLogic.Login(c.Request.Context(), cr)
	response.Response(c, resp, err)
}

func (uh *UserHandler) Register(c *gin.Context) {
	cr := middleware.GetBind[types.RegisterReq](c)
	global.Log.Info(cr)
	resp, err := uh.userLogic.Register(c.Request.Context(), cr)
	response.Response(c, resp, err)
}

func (uh *UserHandler) ResetPassword(c *gin.Context) {
	cr := middleware.GetBind[types.ResetPasswordReq](c)
	global.Log.Info(cr)
	resp, err := uh.userLogic.ResetPassword(c.Request.Context(), cr)
	response.Response(c, resp, err)
}

func (uh *UserHandler) GetLoginCode(c *gin.Context) {
	cr := middleware.GetBind[types.GetCodeReq](c)
	global.Log.Info(cr)
	resp, err := uh.userLogic.GetLoginCode(c.Request.Context(), cr)
	response.Response(c, resp, err)
}

func (uh *UserHandler) GetRegisterCode(c *gin.Context) {
	cr := middleware.GetBind[types.GetCodeReq](c)
	global.Log.Info(cr)
	resp, err := uh.userLogic.GetRegisterCode(c.Request.Context(), cr)
	response.Response(c, resp, err)
}

func (uh *UserHandler) GetResetCode(c *gin.Context) {
	cr := middleware.GetBind[types.GetCodeReq](c)
	global.Log.Info(cr)
	resp, err := uh.userLogic.GetResetCode(c.Request.Context(), cr)
	response.Response(c, resp, err)
}

func (uh *UserHandler) UpdateAvatar(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.Response(c, nil, err)
		return
	}
	resp, err := uh.userLogic.UpdateAvatar(c.Request.Context(), jwtx.GetUserID(c), file)
	if err != nil {
		response.Response(c, nil, err)
	}
	response.Response(c, resp, c.SaveUploadedFile(file, fmt.Sprintf("%s/%s", global.FILE_STORE_PATH, resp)))
}
