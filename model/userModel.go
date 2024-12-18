package model

/**
 * @Author nico
 * @Date 2024-12-18
 * @File: userModel.go
 * @Description:
 */

import (
	"gin-init/model/entity"
)

var _ UserRepoInterface = (*UserRepository)(nil)

type UserRepoInterface interface {
	BaseRepoInterface[entity.TbUser]
	FindByEmail(email string) (*entity.TbUser, error)
}

type UserRepository struct {
	table string
	*BaseRepository[entity.TbUser]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		table:          "tb_user",
		BaseRepository: NewBaseRepository[entity.TbUser](),
	}
}

func (r *UserRepository) FindByEmail(email string) (*entity.TbUser, error) {
	var user entity.TbUser
	if err := r.db.Find(user, email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
