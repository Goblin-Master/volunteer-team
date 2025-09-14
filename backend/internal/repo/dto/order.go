package dto

import (
	"time"
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/infrastructure/model"
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
