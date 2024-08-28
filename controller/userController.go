package controller

import (
	"gin-init/common"
	"gin-init/model/dto"
	"gin-init/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	//
	UserService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (uC *UserController) GetUserList(c *gin.Context) {

	var query dto.QueryPagination
	if err := c.ShouldBindQuery(&query); err != nil {
		common.Failed(c, common.ParamsError)
		return
	}

	// var body dto.UserListFilterDTO
	var body map[string]interface{}
	if err := c.ShouldBindJSON(&body); err != nil {
		common.Failed(c, common.ParamsError)
		return
	}

	//
	data, err := uC.UserService.GetUserList(c, query, body)
	if err != nil {
		common.Failed(c, common.UnKnownError)
		return
	}

	common.SuccessWithData(c, data)
}

func (uC *UserController) GetUserDetail(c *gin.Context) {
	//
	id := c.Query("id")

	// 校验

	// 调用服务层
	data, err := uC.UserService.GetUserDetail(c, id)
	if err != nil {
		common.Failed(c, common.UnKnownError)
		return
	}

	//
	common.SuccessWithData(c, data)
}

func (uC *UserController) CreateUser(context *gin.Context) {

}
