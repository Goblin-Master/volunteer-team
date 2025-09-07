package middleware

import (
	"github.com/gin-gonic/gin"
	"volunteer-team/backend/internal/infrastructure/utils/response"
)

func BindJsonMiddleware[T any](c *gin.Context) {
	var cr T
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.Response(c, nil, err)
		c.Abort()
		return
	}
	c.Set("request", cr)
}
func BindQueryMiddleware[T any](c *gin.Context) {
	var cr T
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.Response(c, nil, err)
		c.Abort()
		return
	}
	c.Set("request", cr)
}
func BindUriMiddleware[T any](c *gin.Context) {
	var cr T
	err := c.ShouldBindUri(&cr)
	if err != nil {
		response.Response(c, nil, err)
		c.Abort()
		return
	}
	c.Set("request", cr)
}
func GetBind[T any](c *gin.Context) T {
	return c.MustGet("request").(T)
}
