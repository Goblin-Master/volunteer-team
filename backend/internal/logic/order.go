package logic

import (
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/infrastructure/model"
	"volunteer-team/backend/internal/infrastructure/pkg/jwtx"
	"volunteer-team/backend/internal/infrastructure/types"
	"volunteer-team/backend/internal/repo"
)

type IOrderLogic interface {
	CreateOrder(int64, types.CreateOrderReq) (string, error)
	GetOrderList(int64, jwtx.Role) (types.OrderListResp, error)
}
type OrderLogic struct {
	orderRepo *repo.OrderRepo
}

func NewOrderLogic() *OrderLogic {
	return &OrderLogic{
		orderRepo: repo.NewOrderRepo(),
	}
}

var _ IOrderLogic = (*OrderLogic)(nil)

func (ol *OrderLogic) CreateOrder(UserID int64, req types.CreateOrderReq) (string, error) {
	err := ol.orderRepo.CreateOrder(UserID, req)
	if err != nil {
		global.Log.Error(err)
		return "", DEFAULT_ERROR
	}
	return "创建订单成功", nil
}

func (ol *OrderLogic) GetOrderList(userID int64, role jwtx.Role) (types.OrderListResp, error) {
	var resp types.OrderListResp
	var list []model.Order
	switch role {
	case jwtx.COMMON_USER: // 2
		// 普通用户：只查自己的单
		data, err := ol.orderRepo.GetOrderListByCommon(userID)
		if err != nil {
			global.Log.Error(err)
			return resp, DEFAULT_ERROR
		}
		list = data
	case jwtx.INTERNAL_USER: // 1
		// 内部人员：查全部“未处理”
		data, err := ol.orderRepo.GetOrderListByInternal()
		if err != nil {
			global.Log.Error(err)
			return resp, DEFAULT_ERROR
		}
		list = data
	default:
		return resp, DEFAULT_ERROR
	}
	var orders []types.OrderItem
	for _, v := range list {
		order := types.OrderItem{
			ID:                 v.ID,
			ProblemDescription: v.ProblemDescription,
			Ctime:              v.Ctime,
		}
		orders = append(orders, order)
	}
	resp.Orders = orders
	return resp, nil
}
