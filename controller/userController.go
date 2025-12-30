package controller

import (
	"gin-init/common/response"
	"gin-init/model/entity"
	"gin-init/model/types"
	"gin-init/service"

	"github.com/gin-gonic/gin"
)

var _ UserControllerInterface = (*UserController)(nil)

type UserControllerInterface interface {
	BaseControllerInterface[entity.TbUser]
	ChangePassword(c *gin.Context)
}

type UserController struct {
	*BaseController[entity.TbUser]
	UserService service.UserServiceInterface
}

func NewUserController() UserControllerInterface {
	return &UserController{
		BaseController: NewBaseController[entity.TbUser](),
		UserService:    service.NewUserService(),
	}
}

func (c *UserController) ChangePassword(ctx *gin.Context) {

}

// 重写 Create
func (c *UserController) Create(ctx *gin.Context) {
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
