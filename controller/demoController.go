package controller

import (
	"gin-init/common"
	"gin-init/service"
	"github.com/gin-gonic/gin"
)

type DemoController struct {
	//
	demoService *service.DemoService
}

func NewDemoController(demoService *service.DemoService) *DemoController {
	return &DemoController{demoService: demoService}
}

func (d *DemoController) SendMsgWithRabbitMQ(c *gin.Context) {
	d.demoService.SendMsgWithRabbitMQ(c)

	//
	common.SuccessWithoutData(c)
}
