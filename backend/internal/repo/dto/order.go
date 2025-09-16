package dto

import (
	"context"
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
func (od *OrderDto) AddOrder(ctx context.Context, order model.Order) error {
	order.Ctime = time.Now().UnixMilli()
	order.Utime = time.Now().UnixMilli()
	return global.DB.WithContext(ctx).Create(&order).Error
}

func (od *OrderDto) GetOrderListByCommon(ctx context.Context, userID int64) ([]model.Order, error) {
	var list []model.Order
	err := global.DB.Model(&model.Order{}).WithContext(ctx).Where("user_id = ?", userID).
		Order("ctime DESC").
		Find(&list).Error
	if err != nil {
		return nil, DEFAULT_ERROR
	}
	return list, nil
}

func (od *OrderDto) GetOrderListByInternal(ctx context.Context) ([]model.Order, error) {
	var list []model.Order
	err := global.DB.WithContext(ctx).Where("state = ?", enum.UNTACKLE).
		Order("ctime ASC"). // 先报先处理
		Find(&list).Error
	if err != nil {
		return nil, DEFAULT_ERROR
	}
	return list, nil
}

func (od *OrderDto) GetOrderByID(ctx context.Context, id int) (model.Order, error) {
	var order model.Order
	err := global.DB.WithContext(ctx).Where("id = ?", id).Take(&order).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return order, ORDER_NOT_EXIST
		}
		global.Log.Error(err)
		return order, DEFAULT_ERROR
	}
	return order, nil
}

func (od *OrderDto) UpdateOrderStateByID(ctx context.Context, id int) error {
	// 直接 WHERE + Update，只改 password 字段
	result := global.DB.Model(&model.Order{}).
		WithContext(ctx).
		Where("id = ?", id).
		Update("state", enum.TACKLE)

	if result.Error != nil {
		global.Log.Error(result.Error)
		return DEFAULT_ERROR
	}
	if result.RowsAffected == 0 {
		return ORDER_NOT_EXIST // 没有匹配行
	}
	return nil
}
