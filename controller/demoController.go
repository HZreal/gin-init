package controller

import (
	"gin-init/common"
	"gin-init/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DemoController struct {
	//
	demoService *service.DemoService
}

func NewDemoController(demoService *service.DemoService) *DemoController {
	return &DemoController{demoService: demoService}
}

func (d *DemoController) SendMsgWithRabbitMQ(c *gin.Context) {
	err := d.demoService.SendMsgWithRabbitMQ(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Failed(common.UnKnownError))
	}
	c.JSON(http.StatusOK, common.SuccessWithoutData())
}
