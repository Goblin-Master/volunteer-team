package handler

import (
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/infrastructure/middleware"
	"volunteer-team/backend/internal/infrastructure/pkg/jwtx"
	"volunteer-team/backend/internal/infrastructure/types"
	"volunteer-team/backend/internal/infrastructure/utils/response"
	"volunteer-team/backend/internal/logic"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderLogic *logic.OrderLogic
}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{
		orderLogic: logic.NewOrderLogic(),
	}
}

func (oh *OrderHandler) CreateOrder(c *gin.Context) {
	cr := middleware.GetBind[types.CreateOrderReq](c)
	global.Log.Info(cr)
	resp, err := oh.orderLogic.CreateOrder(jwtx.GetUserID(c), cr)
	response.Response(c, resp, err)
}

func (oh *OrderHandler) GetOrderList(c *gin.Context) {
	userID, role := jwtx.GetUserID(c), jwtx.GetRole(c)
	global.Log.Info(userID, "  ", role)
	resp, err := oh.orderLogic.GetOrderList(userID, role)
	response.Response(c, resp, err)
}

func (oh *OrderHandler) OrderDetail(c *gin.Context) {
	cr := middleware.GetBind[types.OrderDetailReq](c)
	global.Log.Info(cr)
	//传role是为了防止恶意攻击，如直接调用接口会泄密
	resp, err := oh.orderLogic.OrderDetail(jwtx.GetUserID(c), jwtx.GetRole(c), cr.ID)
	response.Response(c, resp, err)
}
