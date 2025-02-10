package controller

import (
	"gin-init/common/response"
	"gin-init/model/entity"
	"gin-init/model/types"
	"gin-init/service"
	"github.com/gin-gonic/gin"
)

var _ UserControllerInterface = (*UserController2)(nil)

type UserControllerInterface interface {
	BaseControllerInterface[entity.TbUser]
	ChangePassword(c *gin.Context)
}

type UserController2 struct {
	*BaseController[entity.TbUser]
	UserService service.UserServiceInterface
}

func NewUserController2() UserControllerInterface {
	return &UserController2{
		BaseController: NewBaseController[entity.TbUser](),
		UserService:    service.NewUserService2(),
	}
}

func (c *UserController2) ChangePassword(ctx *gin.Context) {

}

// 重写 Create
func (c *UserController2) Create(ctx *gin.Context) {
	var body types.UserCreateDTO

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.Failed(ctx, response.ParamsError)
		return
	}

	user := &entity.TbUser{
		Username: body.Username,
		Password: body.Password,
		Phone:    body.Phone,
		Age:      body.Age,
	}

	//
	err := c.UserService.Create(user)
	if err != nil {
		response.Failed(ctx, response.UnKnownError)
	}

	response.SuccessWithData(ctx, user)
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

func (uC *UserController) GetUserList(c *gin.Context) {

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

func (uC *UserController) GetUserDetail(c *gin.Context) {
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

func (uC *UserController) CreateUser(c *gin.Context) {
	var body types.UserCreateDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		response.Failed(c, response.ParamsError)
		return
	}

	//
	user := uC.UserService.CreateUser(c, body)

	response.SuccessWithData(c, user)
}

func (uC *UserController) UpdateUser(c *gin.Context) {
	var body types.UserUpdateDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		response.Failed(c, response.ParamsError)
		return
	}

	//
	user := uC.UserService.UpdateUser(c, body)

	response.SuccessWithData(c, user)

}

func (uC *UserController) DeleteUser(c *gin.Context) {
	var body types.BodyJsonId

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
