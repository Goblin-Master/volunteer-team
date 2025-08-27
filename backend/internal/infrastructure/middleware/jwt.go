package middleware

import (
	"github.com/gin-gonic/gin"
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/infrastructure/pkg/jwtx"
)

func Authentication(c *gin.Context) {
	UserID, Role, err := jwtx.ParseToken(c)
	if err != nil {
		global.Log.Error(err.Error())
		c.String(401, err.Error())
		c.Abort()
		return
	}
	//将用户id和角色加入ctx
	c.Set(global.TOKEN_USER_ID, UserID)
	c.Set(global.TOKEN_ROLE, Role)
	c.Next()
}
