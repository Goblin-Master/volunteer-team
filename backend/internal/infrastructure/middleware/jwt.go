package middleware

import (
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/infrastructure/pkg/jwtx"
	"volunteer-team/backend/internal/infrastructure/utils/response"

	"github.com/gin-gonic/gin"
)

func Authentication(role jwtx.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		UserID, Role, err := jwtx.ParseToken(c)
		if err != nil {
			global.Log.Error(err.Error())
			c.String(401, err.Error())
			c.Abort()
			return
		}
		if Role < role {
			response.Response(c, nil, jwtx.ErrPermissionDenied)
			c.Abort()
			return
		}
		//将用户id和角色加入ctx
		c.Set(global.TOKEN_USER_ID, UserID)
		c.Set(global.TOKEN_ROLE, Role)
		c.Next()
	}
}
