package dto

import (
	"time"
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/infrastructure/model"
	"volunteer-team/backend/internal/infrastructure/model/enum"
)

type OrderDto struct{}

func NewOrderDto() *OrderDto {
	return &OrderDto{}
}
func (od *OrderDto) AddOrder(order model.Order) error {
	order.Ctime = time.Now().UnixMilli()
	order.Utime = time.Now().UnixMilli()
	return global.DB.Create(&order).Error
}

func (od *OrderDto) GetOrderListByCommon(userID int64) ([]model.Order, error) {
	var list []model.Order
	err := global.DB.Model(&model.Order{}).Where("user_id = ?", userID).
		Order("ctime DESC").
		Find(&list).Error
	if err != nil {
		return nil, DEFAULT_ERROR
	}
	return list, nil
}

func (od *OrderDto) GetOrderListByInternal() ([]model.Order, error) {
	var list []model.Order
	err := global.DB.Where("status = ?", enum.UNTACKLE).
		Order("ctime ASC"). // 先报先处理
		Find(&list).Error
	if err != nil {
		return nil, DEFAULT_ERROR
	}
	return list, nil
}
