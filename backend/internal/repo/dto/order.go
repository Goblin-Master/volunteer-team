package dto

import (
	"errors"
	"time"
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/infrastructure/model"
	"volunteer-team/backend/internal/infrastructure/model/enum"

	"gorm.io/gorm"
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

func (od *OrderDto) GetOrderByID(id int) (model.Order, error) {
	var order model.Order
	err := global.DB.Where("id = ?", id).Take(&order).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return order, ORDER_NOT_EXIST
		}
		global.Log.Error(err)
		return order, DEFAULT_ERROR
	}
	return order, nil
}
