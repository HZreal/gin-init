package service

/**
 * @Author elastic·H
 * @Date 2024-08-27
 * @File: demoService.go
 * @Description:
 */

import (
	"encoding/json"
	"gin-init/service/common"
	"github.com/gin-gonic/gin"
)

type DemoService struct {
	RabbitMQService *common.RabbitMQService
}

func NewDemoService(rabbitMQService *common.RabbitMQService) *DemoService {
	return &DemoService{RabbitMQService: rabbitMQService}
}

func (s *DemoService) SendMsgWithRabbitMQ(c *gin.Context) {
	// mock a msg
	var msg = struct {
		Id   int
		Data string
	}{
		Id:   1,
		Data: "hello world",
	}

	body, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	//
	s.RabbitMQService.SendMQMsq111(body)

}
