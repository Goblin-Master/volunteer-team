package repo

import (
	"volunteer-team/backend/internal/infrastructure/model"
	"volunteer-team/backend/internal/infrastructure/types"
	"volunteer-team/backend/internal/repo/dto"
)

type IOrderRepo interface {
	CreateOrder(int64, types.CreateOrderReq) error
	GetOrderListByCommon(int64) ([]model.Order, error)
	GetOrderListByInternal() ([]model.Order, error)
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
		WechatID:           req.WachatID,
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
