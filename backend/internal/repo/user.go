package repo

import (
	"volunteer-team/backend/internal/infrastructure/model"
	"volunteer-team/backend/internal/repo/dto"
)

type IUserRepo interface {
	LoginWithAccount(account string, password string) (model.User, error)
	LoginWithEmail(email string) (model.User, error)
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
