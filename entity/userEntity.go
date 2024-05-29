package entity

import (
	"gorm.io/gorm"
)

type UserEntity struct {
	gorm.Model

	UserAccount  string  `json:"user_account" gorm:"size:256;comment:用户登录账号;unique" binding:"required,min=4"`
	UserName     *string `json:"user_name,omitempty" gorm:"size:256;comment:用户名"`
	Gender       uint    `json:"gender" gorm:"comment:用户性别;size:2"`
	UserPassword string  `json:"user_password,omitempty" gorm:"size:256;comment:用户密码" binding:"required,min=8"`
	Phone        *string `json:"phone,omitempty" gorm:"size:256;comment:用户手机"`
	Email        *string `json:"email,omitempty" gorm:"size:256;comment:用户邮箱"`
	UserStatus   uint    `json:"user_status"  gorm:"size:2;comment:用户状态 0 表示正常"`
}
