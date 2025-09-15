package dto

import "errors"

var (
	DEFAULT_ERROR = errors.New("默认错误")
)

var (
	USER_NOT_EXIST = errors.New("用户不存在")
)

var (
	ORDER_NOT_EXIST = errors.New("订单不存在")
)
