package dto

import (
	"errors"
	"time"
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/infrastructure/model"

	"gorm.io/gorm"
)

type UserDto struct{}

func NewUserDto() *UserDto {
	return &UserDto{}
}

var (
	USER_NOT_EXIST = errors.New("用户不存在")
	DEFAULT_ERROR  = errors.New("默认错误")
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

func (ud *UserDto) AddUser(user model.User) error {
	user.Ctime = time.Now().UnixMilli()
	user.Utime = time.Now().UnixMilli()
	return global.DB.Create(&user).Error
}

func (ud *UserDto) UpdatePasswordByEmail(email, newPassword string) error {
	// 直接 WHERE + Update，只改 password 字段
	result := global.DB.Model(&model.User{}).
		Where("email = ?", email).
		Update("password", newPassword)

	if result.Error != nil {
		global.Log.Error(result.Error)
		return DEFAULT_ERROR
	}
	if result.RowsAffected == 0 {
		return USER_NOT_EXIST // 没有匹配行
	}
	return nil
}

func (ud *UserDto) UpdateAvatarByID(userID int64, avatar string) error {
	result := global.DB.Model(&model.User{}).
		Where("user_id = ?", userID).
		Update("avatar", avatar)

	if result.Error != nil {
		global.Log.Error(result.Error)
		return DEFAULT_ERROR
	}
	if result.RowsAffected == 0 {
		return USER_NOT_EXIST // 没有匹配行
	}
	return nil
}
