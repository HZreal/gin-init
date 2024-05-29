package controller

import (
	"gin-init/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	//
	userService *service.UserService
}

func (uC *UserController) getUserList(c *gin.Context) {
	//
	body := make(map[string]interface{})
	result := uC.userService.GetUserList(body)

	c.JSON(http.StatusOK, result)

}

func (uC *UserController) getUserDetail(c *gin.Context) {
	//
	id := 1
	result := uC.userService.GetUserDetail(id)

	c.JSON(http.StatusOK, result)
}
