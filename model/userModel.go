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

var _ TbUserModelInterface = (*TbUserModel)(nil)

type TbUserModelInterface interface {
	BaseModelInterface[entity.TbUser]
	FindByEmail(email string) (*entity.TbUser, error)
}

type TbUserModel struct {
	table string
	*BaseModel[entity.TbUser]
}

func NewTbUserModel() *TbUserModel {
	return &TbUserModel{
		table:     "tb_user",
		BaseModel: NewBaseModel[entity.TbUser](),
	}
}

func (r *TbUserModel) FindByEmail(email string) (*entity.TbUser, error) {
	var user entity.TbUser
	if err := r.db.Find(user, email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
