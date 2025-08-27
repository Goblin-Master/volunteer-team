package response

import (
	"github.com/gin-gonic/gin"
)

type Body struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(c *gin.Context, resp interface{}, err error) {
	var body Body
	if err != nil {
		body.Code = -1
		body.Message = err.Error()
		body.Data = nil
	} else {
		body.Code = 0
		body.Message = "OK"
		body.Data = resp
	}
	c.JSON(200, body)
}
