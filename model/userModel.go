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
	BaseRepoInterface[entity.UserModel]
	FindByEmail(email string) (*entity.UserModel, error)
}

type UserRepository struct {
	table string
	*BaseRepository[entity.UserModel]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		table:          "tb_user",
		BaseRepository: NewBaseRepository[entity.UserModel](),
	}
}

func (r *UserRepository) FindByEmail(email string) (*entity.UserModel, error) {
	var user entity.UserModel
	if err := r.db.Find(user, email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
