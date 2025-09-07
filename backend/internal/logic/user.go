package logic

import (
	"errors"
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/infrastructure/pkg/jwtx"
	"volunteer-team/backend/internal/infrastructure/types"
	"volunteer-team/backend/internal/repo"
)

type IUserLogic interface {
	LoginLogic(req types.LoginReq) (types.LoginResp, error)
}
type UserLogic struct {
	userRepo *repo.UserRepo
}

func NewUserLogic() *UserLogic {
	return &UserLogic{
		userRepo: repo.NewUserRepo(),
	}
}

var _ IUserLogic = (*UserLogic)(nil)

var (
	LOGIN_WITH_FAILED_WAY     = errors.New("暂不支持这种登录方式")
	ACCOUNT_OR_PASSWORD_ERROR = errors.New("账号或密码错误")
	EMAIL_ERROR               = errors.New("邮箱错误")
	DEFAULT_ERROR             = errors.New("默认错误")
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
