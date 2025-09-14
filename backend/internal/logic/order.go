package logic

import (
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/infrastructure/types"
	"volunteer-team/backend/internal/repo"
)

type IOrderLogic interface {
	CreateOrder(int64, types.CreateOrderReq) (string, error)
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
