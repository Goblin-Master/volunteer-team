package repo

import (
	"context"
	"errors"
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/infrastructure/model"
	"volunteer-team/backend/internal/infrastructure/types"
	"volunteer-team/backend/internal/repo/dto"
)

type IOrderRepo interface {
	CreateOrder(ctx context.Context, userID int64, orderID int64, req types.CreateOrderReq) error
	GetOrderListByCommon(ctx context.Context, userID int64) ([]model.Order, error)
	GetOrderListByInternal(ctx context.Context) ([]model.Order, error)
	GetOrderDetail(ctx context.Context, orderID int64) (model.Order, error)
	UpdateOrderState(ctx context.Context, orderID int64) error
}
type OrderRepo struct {
	orderDto *dto.OrderDto
}

func NewOrderRepo() *OrderRepo {
	return &OrderRepo{
		orderDto: dto.NewOrderDto(),
	}
}

var _ IOrderRepo = (*OrderRepo)(nil)

func (or *OrderRepo) CreateOrder(ctx context.Context, userID int64, orderID int64, req types.CreateOrderReq) error {
	order := model.Order{
		OrderID:            orderID,
		UserID:             userID,
		StudentID:          req.StudentID,
		Username:           req.Username,
		CampusLocation:     req.CampusLocation,
		Department:         req.Department,
		PhoneNumber:        req.PhoneNumber,
		WechatID:           req.WechatID,
		Address:            req.Address,
		DeviceModel:        req.DeviceModel,
		OSVersion:          req.OSVersion,
		ProblemDescription: req.ProblemDescription,
		Notes:              req.Notes,
	}
	return or.orderDto.AddOrder(ctx, order)
}

func (or *OrderRepo) GetOrderListByCommon(ctx context.Context, userID int64) ([]model.Order, error) {
	return or.orderDto.GetOrderListByCommon(ctx, userID)
}

func (or *OrderRepo) GetOrderListByInternal(ctx context.Context) ([]model.Order, error) {
	return or.orderDto.GetOrderListByInternal(ctx)
}

func (or *OrderRepo) GetOrderDetail(ctx context.Context, orderID int64) (model.Order, error) {
	data, err := or.orderDto.GetOrderDetailByID(ctx, orderID)
	if err != nil {
		if errors.Is(err, dto.ORDER_NOT_EXIST) {
			return data, ORDER_NOT_EXIST
		}
		global.Log.Info(err)
		return data, DEFAULT_ERROR
	}
	return data, nil
}

func (or *OrderRepo) UpdateOrderState(ctx context.Context, orderID int64) error {
	err := or.orderDto.UpdateOrderStateByID(ctx, orderID)
	if err != nil {
		if errors.Is(err, dto.ORDER_NOT_EXIST) {
			return ORDER_NOT_EXIST
		}
		global.Log.Error(err)
		return DEFAULT_ERROR
	}
	return nil
}
