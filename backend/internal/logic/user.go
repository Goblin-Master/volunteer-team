package logic

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/infrastructure/pkg/emailx"
	"volunteer-team/backend/internal/infrastructure/pkg/jwtx"
	"volunteer-team/backend/internal/infrastructure/types"
	"volunteer-team/backend/internal/infrastructure/utils/code"
	"volunteer-team/backend/internal/infrastructure/utils/fileUtils"
	"volunteer-team/backend/internal/infrastructure/utils/snowflake"
	"volunteer-team/backend/internal/repo"
)

type IUserLogic interface {
	Login(types.LoginReq) (types.LoginResp, error)
	Register(types.RegisterReq) (types.RegisterResp, error)
	GetLoginCode(types.GetCodeReq) (types.GetCodeResp, error)
	GetRegisterCode(types.GetCodeReq) (types.GetCodeResp, error)
	GetResetCode(types.GetCodeReq) (types.GetCodeResp, error)
	ResetPassword(types.ResetPasswordReq) (types.ResetPasswordResp, error)
	UpdateAvatar(int64, *multipart.FileHeader) (string, error)
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

func (ul *UserLogic) Login(req types.LoginReq) (types.LoginResp, error) {
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

func (ul *UserLogic) Register(req types.RegisterReq) (types.RegisterResp, error) {
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

func (ul *UserLogic) ResetPassword(req types.ResetPasswordReq) (types.ResetPasswordResp, error) {
	var resp types.ResetPasswordResp
	if ok := ul.email.VerifyCode(req.Email, req.Code); !ok {
		return resp, CODE_VERIFY_ERROR
	}
	err := ul.userRepo.ResetPassword(req.Email, req.NewPassword)
	if err != nil {
		if errors.Is(err, repo.USER_NOT_EXIST) {
			return resp, USER_NOT_EXIST
		}
		global.Log.Error(err)
		return resp, DEFAULT_ERROR
	}
	resp.Message = "重置密码成功！"
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

func (ul *UserLogic) UpdateAvatar(userID int64, file *multipart.FileHeader) (string, error) {
	//图片大小限制
	if file.Size > global.FILE_MAX_SIZE {
		return "", FILE_OVER_SIZE
	}
	//图片格式限制
	filename := file.Filename
	suffix, err := fileUtils.FileSuffixJudge(filename)
	if err != nil {
		global.Log.Error(err)
		//FileSuffixJudge内部已经做好err处理
		return "", err
	}
	//文件hash（不同文件名，内容相同视为同一个文件）
	data, err := file.Open()
	defer func() {
		if e := data.Close(); e != nil {
			global.Log.Error(e)
		}
	}()
	if err != nil {
		global.Log.Error(err)
		return "", FILE_READ_ERROR
	}
	byteData, _ := io.ReadAll(data)
	hash := fileUtils.Md5(byteData)
	// 更新用户头像
	filePath := fmt.Sprintf("%s.%s", hash, suffix)
	err = ul.userRepo.UpdateAvatarByID(userID, filePath)
	if err != nil {
		if errors.Is(err, repo.USER_NOT_EXIST) {
			return "", USER_NOT_EXIST
		}
		global.Log.Error(err)
		return "", DEFAULT_ERROR
	}
	return filePath, nil
}
