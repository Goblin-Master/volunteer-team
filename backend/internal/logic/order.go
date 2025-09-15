package logic

import (
	"errors"
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/infrastructure/model"
	"volunteer-team/backend/internal/infrastructure/pkg/jwtx"
	"volunteer-team/backend/internal/infrastructure/types"
	"volunteer-team/backend/internal/repo"
)

type IOrderLogic interface {
	CreateOrder(int64, types.CreateOrderReq) (string, error)
	GetOrderList(int64, jwtx.Role) (types.OrderListResp, error)
	OrderDetail(int64, jwtx.Role, int) (types.OrderDetailResp, error)
	FinishOrder(jwtx.Role, int) (string, error)
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

func (ol *OrderLogic) OrderDetail(userID int64, role jwtx.Role, id int) (types.OrderDetailResp, error) {
	var resp types.OrderDetailResp
	data, err := ol.orderRepo.OrderDetail(id)
	if err != nil {
		global.Log.Error(err)
		return resp, DEFAULT_ERROR
	}
	//防止非内部人员查看到别人的订单
	if role != jwtx.INTERNAL_USER && data.UserID != userID {
		return resp, ORDER_IS_FORBIDDEN
	}
	//组装数据
	resp.Notes = data.Notes
	resp.OSVersion = data.OSVersion
	resp.Ctime = data.Ctime
	resp.ProblemDescription = data.ProblemDescription
	resp.WachatID = data.WechatID
	resp.DeviceModel = data.DeviceModel
	resp.PhoneNumber = data.PhoneNumber
	resp.Username = data.Username
	resp.Department = data.Department
	resp.StudentID = data.StudentID
	resp.CampusLocation = data.CampusLocation
	resp.Address = data.Address
	return resp, nil
}

func (ol *OrderLogic) FinishOrder(role jwtx.Role, id int) (string, error) {
	if role != jwtx.INTERNAL_USER {
		return "", ORDER_IS_FORBIDDEN
	}
	err := ol.orderRepo.UpdateOrderState(id)
	if err != nil {
		if errors.Is(err, repo.ORDER_NOT_EXIST) {
			return "", ORDER_NOT_EXIST
		}
		global.Log.Error(err)
		return "", DEFAULT_ERROR
	}
	return "订单已完成", nil
}
