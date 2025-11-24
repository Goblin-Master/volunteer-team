package dto

import "errors"

var (
	ErrDefault = errors.New("默认错误")
)

var (
	ErrUserNotExist = errors.New("用户不存在")
)

var (
	ErrOrderNotExist = errors.New("订单不存在")
)

var (
	ErrSummaryNotExist = errors.New("修机总结不存在")
)
