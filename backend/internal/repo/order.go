package repo

import (
	"errors"
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/infrastructure/model"
	"volunteer-team/backend/internal/infrastructure/types"
	"volunteer-team/backend/internal/repo/dto"
)

type IOrderRepo interface {
	CreateOrder(userID int64, req types.CreateOrderReq) error
	GetOrderListByCommon(userID int64) ([]model.Order, error)
	GetOrderListByInternal() ([]model.Order, error)
	OrderDetail(id int) (model.Order, error)
	UpdateOrderState(id int) error
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

func (or *OrderRepo) CreateOrder(userID int64, req types.CreateOrderReq) error {
	order := model.Order{
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
	return or.orderDto.AddOrder(order)
}

func (or *OrderRepo) GetOrderListByCommon(userID int64) ([]model.Order, error) {
	return or.orderDto.GetOrderListByCommon(userID)
}

func (or *OrderRepo) GetOrderListByInternal() ([]model.Order, error) {
	return or.orderDto.GetOrderListByInternal()
}

func (or *OrderRepo) OrderDetail(id int) (model.Order, error) {
	return or.orderDto.GetOrderByID(id)
}

func (or *OrderRepo) UpdateOrderState(id int) error {
	err := or.orderDto.UpdateOrderStateByID(id)
	if err != nil {
		if errors.Is(err, dto.ORDER_NOT_EXIST) {
			return ORDER_NOT_EXIST
		}
		global.Log.Error(err)
		return DEFAULT_ERROR
	}
	return nil
}
