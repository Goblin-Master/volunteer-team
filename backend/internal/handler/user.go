package handler

import (
	"github.com/gin-gonic/gin"
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/infrastructure/middleware"
	"volunteer-team/backend/internal/infrastructure/types"
	"volunteer-team/backend/internal/infrastructure/utils/response"
	"volunteer-team/backend/internal/logic"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}
func (u *UserHandler) LoginHandler(c *gin.Context) {
	cr := middleware.GetBind[types.LoginReq](c)
	global.Log.Info(cr)
	resp, err := logic.NewUserLogic().LoginLogic(cr)
	response.Response(c, resp, err)
}
