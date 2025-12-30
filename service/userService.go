package service

import (
	"errors"
	"gin-init/model"
	"gin-init/model/entity"
	"gin-init/service/common"
)

var _ UserServiceInterface = (*UserService)(nil)

type UserServiceInterface interface {
	BaseServiceInterface[entity.TbUser]
	ChangePassword(id uint, oldPassword, newPassword string) error
}

type UserService struct {
	*BaseService[entity.TbUser]
	TbUserModel  model.TbUserModelInterface
	RedisService *common.RedisService
}

func NewUserService() UserServiceInterface {
	return &UserService{
		RedisService: common.NewRedisService(),
		TbUserModel:  model.NewTbUserModel(),
		BaseService:  NewBaseService[entity.TbUser](),
	}
}

func (s *UserService) ChangePassword(id uint, oldPassword, newPassword string) error {
	user, err := s.TbUserModel.FindByID(id)
	if err != nil {
		return err
	}
	if user.Password != oldPassword {
		return errors.New("old password is incorrect")
	}
	user.Password = newPassword
	return s.TbUserModel.Update(user)
}
