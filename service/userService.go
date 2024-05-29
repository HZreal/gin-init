package service

import (
	"gin-init/common"
)

type UserService struct {
	// userRepo *entity.UserEntity
}

func (uS *UserService) GetUserList(condition map[string]interface{}) common.Response {
	//
	return common.Response{Code: 0, Msg: "success", Data: []map[string]interface{}{}}

}

func (uS *UserService) GetUserDetail(id int) common.Response {
	//
	// user:= uS.userRepo.
	return common.Response{Code: 0, Msg: "success", Data: make(map[string]interface{})}

}
