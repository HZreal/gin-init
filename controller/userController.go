package controller

import (
	"gin-init/common/response"
	"gin-init/model/dto"
	"gin-init/model/entity"
	"gin-init/service"
	"github.com/gin-gonic/gin"
)

var _ UserControllerInterface = (*UserController2)(nil)

type UserControllerInterface interface {
	BaseControllerInterface[entity.UserModel]
	ChangePassword(c *gin.Context)
}

type UserController2 struct {
	*BaseController[entity.UserModel]
	UserService service.UserServiceInterface
}

func NewUserController2() UserControllerInterface {
	return &UserController2{
		BaseController: NewBaseController[entity.UserModel](),
		UserService:    service.NewUserService2(),
	}
}

func (c *UserController2) ChangePassword(ctx *gin.Context) {

}

type UserController struct {
	//
	UserService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (uC *UserController) GetAllUser(c *gin.Context) {
	//
	var body dto.UsersFilterDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		response.Failed(c, response.ParamsError)
		return
	}

	// 调用服务层
	data := uC.UserService.GetAllUser(c, body)

	//
	response.SuccessWithData(c, data)
}

func (uC *UserController) GetUserList(c *gin.Context) {

	var query dto.QueryPagination
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

func (uC *UserController) GetUserDetail(c *gin.Context) {
	//
	var body dto.QueryId

	if err := c.ShouldBindQuery(&body); err != nil {
		response.Failed(c, response.ParamsError)
		return
	}

	// 调用服务层
	data := uC.UserService.GetUserDetail(c, body.Id)

	//
	response.SuccessWithData(c, data)
}

func (uC *UserController) CreateUser(c *gin.Context) {
	var body dto.UserCreateDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		response.Failed(c, response.ParamsError)
		return
	}

	//
	user := uC.UserService.CreateUser(c, body)

	response.SuccessWithData(c, user)
}

func (uC *UserController) UpdateUser(c *gin.Context) {
	var body dto.UserUpdateDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		response.Failed(c, response.ParamsError)
		return
	}

	//
	user := uC.UserService.UpdateUser(c, body)

	response.SuccessWithData(c, user)

}

func (uC *UserController) DeleteUser(c *gin.Context) {
	var body dto.BodyJsonId

	if err := c.ShouldBindJSON(&body); err != nil {
		response.Failed(c, response.ParamsError)
		return
	}

	//
	uC.UserService.DeleteUser(c, body.Id)

	response.SuccessWithoutData(c)
}

func (uC *UserController) UpdateUserPassword(c *gin.Context) {

}

func (uC *UserController) ResetUserPassword(c *gin.Context) {

}
