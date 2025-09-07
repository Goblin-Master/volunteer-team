package dto

import (
	"errors"
	"gorm.io/gorm"
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/infrastructure/model"
)

type UserDto struct{}

func NewUserDto() *UserDto {
	return &UserDto{}
}

var (
	USER_NOT_EXIST = errors.New("用户不存在")
	DEFAULT_ERROR  = errors.New("默认失败")
)

func (ud *UserDto) VerifyUserByAccount(account, password string) (model.User, error) {
	var user model.User
	err := global.DB.Model(&model.User{}).Where("account = ? and password = ?", account, password).Take(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, USER_NOT_EXIST
		}
		global.Log.Error(err)
		return user, DEFAULT_ERROR
	}
	return user, nil
}
func (ud *UserDto) VerifyUserByEmail(email string) (model.User, error) {
	var user model.User
	err := global.DB.Model(&model.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, USER_NOT_EXIST
		}
		global.Log.Error(err)
		return user, DEFAULT_ERROR
	}
	return user, nil
}
