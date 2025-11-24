package logic

import (
	"context"
	"errors"
	"strconv"
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/infrastructure/model"
	"volunteer-team/backend/internal/infrastructure/pkg/jwtx"
	"volunteer-team/backend/internal/infrastructure/types"
	"volunteer-team/backend/internal/infrastructure/utils/snowflake"
	"volunteer-team/backend/internal/repo"
)

type IOrderLogic interface {
	CreateOrder(ctx context.Context, userID int64, req types.CreateOrderReq) (string, error)
	GetOrderList(ctx context.Context, userID int64, role jwtx.Role) (types.OrderListResp, error)
	GetOrderDetail(ctx context.Context, userID int64, role jwtx.Role, req types.OrderDetailReq) (types.OrderDetailResp, error)
	FinishOrder(ctx context.Context, req types.FinishOrderReq) (string, error)
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

func (ol *OrderLogic) CreateOrder(ctx context.Context, userID int64, req types.CreateOrderReq) (string, error) {
	err := ol.orderRepo.CreateOrder(ctx, userID, snowflake.GetIntID(global.Node), req)
	if err != nil {
		global.Log.Error(err)
		return "", ErrDefault
	}
	return "创建订单成功", nil
}

func (ol *OrderLogic) GetOrderList(ctx context.Context, userID int64, role jwtx.Role) (types.OrderListResp, error) {
	var resp types.OrderListResp
	var list []model.Order
	switch role {
	case jwtx.COMMON_USER:
		// 普通用户：只查自己的单
		data, err := ol.orderRepo.GetOrderListByCommon(ctx, userID)
		if err != nil {
			global.Log.Error(err)
			return resp, ErrDefault
		}
		list = data
	case jwtx.INTERNAL_USER, jwtx.ADMIN:
		// 内部人员：查全部“未处理”
		data, err := ol.orderRepo.GetOrderListByInternal(ctx)
		if err != nil {
			global.Log.Error(err)
			return resp, ErrDefault
		}
		list = data
	default:
		return resp, ErrDefault
	}
	var orders []types.OrderItem
	for _, v := range list {
		order := types.OrderItem{
			OrderID:            v.OrderID,
			ProblemDescription: v.ProblemDescription,
			Ctime:              v.Ctime,
		}
		orders = append(orders, order)
	}
	resp.Orders = orders
	return resp, nil
}

func (ol *OrderLogic) GetOrderDetail(ctx context.Context, userID int64, role jwtx.Role, req types.OrderDetailReq) (types.OrderDetailResp, error) {
	var resp types.OrderDetailResp
	orderID, err := strconv.ParseInt(req.OrderID, 10, 64)
	if err != nil {
		return resp, ErrParamsType
	}
	data, err := ol.orderRepo.GetOrderDetail(ctx, orderID)
	if err != nil {
		if errors.Is(err, repo.ErrOrderNotExist) {
			return resp, ErrOrderNotExist
		}
		global.Log.Error(err)
		return resp, ErrDefault
	}
	//防止非内部人员查看到别人的订单
	if role != jwtx.INTERNAL_USER && data.UserID != userID {
		return resp, ErrOrderIsForbidden
	}
	//组装数据
	resp.Notes = data.Notes
	resp.OSVersion = data.OSVersion
	resp.Ctime = data.Ctime
	resp.ProblemDescription = data.ProblemDescription
	resp.WechatID = data.WechatID
	resp.DeviceModel = data.DeviceModel
	resp.PhoneNumber = data.PhoneNumber
	resp.Username = data.Username
	resp.Department = data.Department
	resp.StudentID = data.StudentID
	resp.CampusLocation = data.CampusLocation
	resp.Address = data.Address
	return resp, nil
}

func (ol *OrderLogic) FinishOrder(ctx context.Context, req types.FinishOrderReq) (string, error) {
	orderID, err := strconv.ParseInt(req.OrderID, 10, 64)
	if err != nil {
		return "", ErrParamsType
	}
	err = ol.orderRepo.UpdateOrderState(ctx, orderID)
	if err != nil {
		if errors.Is(err, repo.ErrOrderNotExist) {
			return "", ErrOrderNotExist
		}
		global.Log.Error(err)
		return "", ErrDefault
	}
	return "订单已完成", nil
}
