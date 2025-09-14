package repo

import (
	"errors"
	"strings"
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/infrastructure/model"
	"volunteer-team/backend/internal/repo/dto"
)

type IUserRepo interface {
	LoginWithAccount(account string, password string) (model.User, error)
	LoginWithEmail(email string) (model.User, error)
	Register(userID int64, username, email, account, password string) error //这个结构默认都注册普通用户
	ResetPassword(email, newPassword string) error
	UpdateAvatarByID(userID int64, url string) error
}
type UserRepo struct {
	userDto *dto.UserDto
}

func NewUserRepo() *UserRepo {
	return &UserRepo{
		userDto: dto.NewUserDto(),
	}
}

var _ IUserRepo = (*UserRepo)(nil)

func (ur *UserRepo) LoginWithAccount(account string, password string) (model.User, error) {
	return ur.userDto.VerifyUserByAccount(account, password)
}

func (ur *UserRepo) LoginWithEmail(email string) (model.User, error) {
	return ur.userDto.VerifyUserByEmail(email)
}

func (ur *UserRepo) Register(userID int64, username, email, account, password string) error {
	user := model.User{
		Username: username,
		UserID:   userID,
		Account:  account,
		Password: password,
		Email:    email,
	}
	err := ur.userDto.AddUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "for key 'user.uni_user_account'") {
			return ACCOUNT_IS_USED
		} else if strings.Contains(err.Error(), "for key 'user.uni_user_email'") {
			return EMAIL_IS_USED
		} else {
			global.Log.Error(err)
			return err
		}
	}
	return nil
}
func (ur *UserRepo) ResetPassword(email, newPassword string) error {
	err := ur.userDto.UpdatePasswordByEmail(email, newPassword)
	if err != nil {
		if errors.Is(err, dto.USER_NOT_EXIST) {
			return USER_NOT_EXIST
		}
		global.Log.Error(err)
		return DEFAULT_ERROR
	}
	return nil
}

func (ur *UserRepo) UpdateAvatarByID(userID int64, url string) error {
	err := ur.userDto.UpdateAvatarByID(userID, url)
	if err != nil {
		if errors.Is(err, dto.USER_NOT_EXIST) {
			return USER_NOT_EXIST
		}
		global.Log.Error(err)
		return DEFAULT_ERROR
	}
	return nil
}
