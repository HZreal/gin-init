package controller

/**
 * @Author nico
 * @Date 2025-12-30
 * @File: demoUserController.go
 * @Description:
 */

import (
	"gin-init/common/response"
	"gin-init/model/types"
	"gin-init/service"

	"github.com/gin-gonic/gin"
)

type DemoUserController struct {
	//
	UserService *service.DemoUserService
}

func NewDemoUserController(userService *service.DemoUserService) *DemoUserController {
	return &DemoUserController{UserService: userService}
}

func (uC *DemoUserController) GetAllUser(c *gin.Context) {
	//
	var body types.UsersFilterDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		response.Failed(c, response.ParamsError)
		return
	}

	// 调用服务层
	data := uC.UserService.GetAllUser(c, body)

	//
	response.SuccessWithData(c, data)
}

func (uC *DemoUserController) GetUserList(c *gin.Context) {

	var query types.QueryPagination
	if err := c.ShouldBindQuery(&query); err != nil {
		response.Failed(c, response.ParamsError)
		return
	}

	// var body dto.UserListFilterDTO
	var body map[string]interface{}
	if err := c.ShouldBindJSON(&body); err != nil {
		response.Failed(c, response.ParamsError)
		return
	}

	//
	data := uC.UserService.GetUserList(c, query, body)

	response.SuccessWithData(c, data)
}

func (uC *DemoUserController) GetUserDetail(c *gin.Context) {
	//
	var body types.QueryId

	if err := c.ShouldBindQuery(&body); err != nil {
		response.Failed(c, response.ParamsError)
		return
	}

	// 调用服务层
	data := uC.UserService.GetUserDetail(c, body.Id)

	//
	response.SuccessWithData(c, data)
}

func (uC *DemoUserController) CreateUser(c *gin.Context) {
	var body types.UserCreateDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		response.Failed(c, response.ParamsError)
		return
	}

	//
	user := uC.UserService.CreateUser(c, body)

	response.SuccessWithData(c, user)
}

func (uC *DemoUserController) UpdateUser(c *gin.Context) {
	var body types.UserUpdateDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		response.Failed(c, response.ParamsError)
		return
	}

	//
	user := uC.UserService.UpdateUser(c, body)

	response.SuccessWithData(c, user)

}

func (uC *DemoUserController) DeleteUser(c *gin.Context) {
	var body types.BodyJsonId

	if err := c.ShouldBindJSON(&body); err != nil {
		response.Failed(c, response.ParamsError)
		return
	}

	//
	uC.UserService.DeleteUser(c, body.Id)

	response.SuccessWithoutData(c)
}

func (uC *DemoUserController) UpdateUserPassword(c *gin.Context) {

}

func (uC *DemoUserController) ResetUserPassword(c *gin.Context) {

}
