package logic

import (
	"errors"
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/infrastructure/pkg/emailx"
	"volunteer-team/backend/internal/infrastructure/pkg/jwtx"
	"volunteer-team/backend/internal/infrastructure/types"
	"volunteer-team/backend/internal/infrastructure/utils/code"
	"volunteer-team/backend/internal/infrastructure/utils/snowflake"
	"volunteer-team/backend/internal/repo"
)

type IUserLogic interface {
	LoginLogic(req types.LoginReq) (types.LoginResp, error)
	RegisterLogic(req types.RegisterReq) (types.RegisterResp, error)
	GetLoginCode(req types.GetCodeReq) (types.GetCodeResp, error)
	GetRegisterCode(req types.GetCodeReq) (types.GetCodeResp, error)
	GetResetCode(req types.GetCodeReq) (types.GetCodeResp, error)
}
type UserLogic struct {
	userRepo *repo.UserRepo
	email    *emailx.EmailX
}

func NewUserLogic() *UserLogic {
	return &UserLogic{
		userRepo: repo.NewUserRepo(),
		email:    emailx.NewEmailX(),
	}
}

var _ IUserLogic = (*UserLogic)(nil)

var (
	LOGIN_WITH_FAILED_WAY     = errors.New("暂不支持这种登录方式")
	ACCOUNT_OR_PASSWORD_ERROR = errors.New("账号或密码错误")
	EMAIL_ERROR               = errors.New("邮箱错误")
	DEFAULT_ERROR             = errors.New("默认错误")
	CODE_GET_ERROR            = errors.New("code获取失败")
	CODE_VERIFY_ERROR         = errors.New("验证码错误")
	EMAIL_IS_USED             = errors.New("邮箱已经被使用")
	ACCOUNT_IS_USED           = errors.New("账号已经被使用")
)

func (ul *UserLogic) LoginLogic(req types.LoginReq) (types.LoginResp, error) {
	var resp types.LoginResp
	switch req.LoginType {
	case global.LOGIN_WITH_ACCOUNT:
		data, err := ul.userRepo.LoginWithAccount(req.Account, req.Password)
		if err != nil {
			return resp, ACCOUNT_OR_PASSWORD_ERROR
		}
		token, err := jwtx.GenToken(jwtx.Claims{
			UserID: data.UserID,
			Role:   jwtx.Role(data.Role),
		})
		if err != nil {
			global.Log.Error(err)
			return resp, DEFAULT_ERROR
		}
		resp.Username = data.Username
		resp.Avatar = data.Avatar
		resp.Token = token
		return resp, nil
	case global.LOGIN_WITH_EMAIL:
		if ok := ul.email.VerifyCode(req.Email, req.Code); !ok {
			return resp, CODE_VERIFY_ERROR
		}
		data, err := ul.userRepo.LoginWithEmail(req.Email)
		if err != nil {
			return resp, EMAIL_ERROR
		}
		token, err := jwtx.GenToken(jwtx.Claims{
			UserID: data.UserID,
			Role:   jwtx.Role(data.Role),
		})
		if err != nil {
			global.Log.Error(err)
			return resp, DEFAULT_ERROR
		}
		resp.Username = data.Username
		resp.Avatar = data.Avatar
		resp.Token = token
		return resp, nil
	default:
		global.Log.Warnf("错误的登录方式:%s", req.LoginType)
		return resp, LOGIN_WITH_FAILED_WAY
	}
}

func (ul *UserLogic) RegisterLogic(req types.RegisterReq) (types.RegisterResp, error) {
	var resp types.RegisterResp
	if ok := ul.email.VerifyCode(req.Email, req.Code); !ok {
		return resp, CODE_VERIFY_ERROR
	}
	err := ul.userRepo.Register(snowflake.GetIntID(global.Node), req.Username, req.Email, req.Account, req.Password)
	if err != nil {
		if errors.Is(err, repo.EMAIL_IS_USED) {
			return resp, EMAIL_IS_USED
		} else if errors.Is(err, repo.ACCOUNT_IS_USED) {
			return resp, ACCOUNT_IS_USED
		} else {
			global.Log.Error(err)
			return resp, DEFAULT_ERROR
		}
	}
	resp.Message = "用户注册成功！"
	return resp, nil
}

func (ul *UserLogic) GetLoginCode(req types.GetCodeReq) (types.GetCodeResp, error) {
	var resp types.GetCodeResp
	c := code.GenCode()
	err := ul.email.SendLoginCode(req.Email, c)
	if err != nil {
		global.Log.Error(err)
		return resp, CODE_GET_ERROR
	}
	resp.Code = c
	return resp, nil
}

func (ul *UserLogic) GetRegisterCode(req types.GetCodeReq) (types.GetCodeResp, error) {
	var resp types.GetCodeResp
	c := code.GenCode()
	err := ul.email.SendBindEmail(req.Email, c)
	if err != nil {
		global.Log.Error(err)
		return resp, CODE_GET_ERROR
	}
	resp.Code = c
	return resp, nil
}

func (ul *UserLogic) GetResetCode(req types.GetCodeReq) (types.GetCodeResp, error) {
	var resp types.GetCodeResp
	c := code.GenCode()
	err := ul.email.SendResetPwdCode(req.Email, c)
	if err != nil {
		global.Log.Error(err)
		return resp, CODE_GET_ERROR
	}
	resp.Code = c
	return resp, nil
}
