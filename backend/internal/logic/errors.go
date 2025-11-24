package logic

import (
	"errors"
	"fmt"
	"volunteer-team/backend/internal/infrastructure/global"
)

var (
	ErrParamsType   = errors.New("参数格式错误")
	ErrDefault      = errors.New("默认错误")
	ErrFileOverSize = fmt.Errorf("文件大小不能超过%dMB", global.FILE_MAX_SIZE/1024/1024)
	ErrFileRead     = errors.New("文件读取失败")
)
var (
	ErrLoginWithFailedWay = errors.New("暂不支持这种登录方式")
	ErrAccountOrPassword  = errors.New("账号或密码错误")
	ErrEmail              = errors.New("邮箱错误")
	ErrCodeGet            = errors.New("code获取失败")
	ErrCodeVerify         = errors.New("验证码错误")
	ErrEmailIsUsed        = errors.New("邮箱已经被使用")
	ErrAccountIsUsed      = errors.New("账号已经被使用")
	ErrUserNotExist       = errors.New("用户不存在")
)

var (
	ErrOrderIsForbidden = errors.New("禁止操作别人的订单")
	ErrOrderNotExist    = errors.New("订单不存在")
)

var (
	ErrSummaryNotExist = errors.New("修机总结不存在")
)
