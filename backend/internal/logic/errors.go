package logic

import (
	"errors"
	"fmt"
	"volunteer-team/backend/internal/infrastructure/global"
)

var (
	DEFAULT_ERROR = errors.New("默认错误")
)
var (
	LOGIN_WITH_FAILED_WAY     = errors.New("暂不支持这种登录方式")
	ACCOUNT_OR_PASSWORD_ERROR = errors.New("账号或密码错误")
	EMAIL_ERROR               = errors.New("邮箱错误")
	CODE_GET_ERROR            = errors.New("code获取失败")
	CODE_VERIFY_ERROR         = errors.New("验证码错误")
	EMAIL_IS_USED             = errors.New("邮箱已经被使用")
	ACCOUNT_IS_USED           = errors.New("账号已经被使用")
	USER_NOT_EXIST            = errors.New("用户不存在")
	FILE_OVER_SIZE            = errors.New(fmt.Sprintf("文件大小不能超过%dMB", global.FILE_MAX_SIZE/1024/1024))
	FILE_READ_ERROR           = errors.New("文件读取失败")
)
