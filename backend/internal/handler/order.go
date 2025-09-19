package handler

import (
	"strconv"
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
	resp, err := oh.orderLogic.CreateOrder(c.Request.Context(), jwtx.GetUserID(c), cr)
	response.Response(c, resp, err)
}

func (oh *OrderHandler) GetOrderList(c *gin.Context) {
	userID, role := jwtx.GetUserID(c), jwtx.GetRole(c)
	global.Log.Info(userID, "  ", role)
	resp, err := oh.orderLogic.GetOrderList(c.Request.Context(), userID, role)
	response.Response(c, resp, err)
}

func (oh *OrderHandler) GetOrderDetail(c *gin.Context) {
	cr := middleware.GetBind[types.OrderDetailReq](c)
	global.Log.Info(cr)
	orderID, err := strconv.ParseInt(cr.OrderID, 10, 64)
	if err != nil {
		response.Response(c, nil, PARAMS_TYPE_ERROR)
		return
	}
	resp, err := oh.orderLogic.GetOrderDetail(c.Request.Context(), jwtx.GetUserID(c), jwtx.GetRole(c), orderID)
	response.Response(c, resp, err)
}

func (oh *OrderHandler) FinishOrder(c *gin.Context) {
	cr := middleware.GetBind[types.FinishOrderReq](c)
	global.Log.Info(cr)
	orderID, err := strconv.ParseInt(cr.OrderID, 10, 64)
	if err != nil {
		response.Response(c, nil, PARAMS_TYPE_ERROR)
		return
	}
	resp, err := oh.orderLogic.FinishOrder(c.Request.Context(), orderID)
	response.Response(c, resp, err)
}
