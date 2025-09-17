package repo

import "errors"

var (
	DEFAULT_ERROR = errors.New("默认错误")
)

var (
	EMAIL_IS_USED   = errors.New("邮箱已经被使用")
	ACCOUNT_IS_USED = errors.New("账号已经被使用")
	USER_NOT_EXIST  = errors.New("用户不存在")
)
var (
	ORDER_NOT_EXIST = errors.New("订单不存在")
)

var (
	SUMMARY_NOT_EXIST = errors.New("总结不存在")
)
