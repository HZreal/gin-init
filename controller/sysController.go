package controller

import (
	"gin-init/common"
	"gin-init/model/dto"
	"gin-init/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SysController struct {
	//
	sysService *service.SysService
}

func (uC *SysController) Login(c *gin.Context) {
	//
	var loginData dto.LoginData
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		// return
	}

	//
	data, err := uC.sysService.Login(c, loginData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Failed(common.UKnownError))
	}

	c.JSON(http.StatusOK, common.SuccessWithData(data))

}

func (uC *SysController) Logout(c *gin.Context) {
	//
	c.JSON(http.StatusOK, gin.H{})

}
