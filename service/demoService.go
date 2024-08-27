package service

/**
 * @Author elastic·H
 * @Date 2024-08-27
 * @File: demoService.go
 * @Description:
 */

import (
	"encoding/json"
	"fmt"
	"gin-init/mq"
	"gin-init/mq/rabbitMQ"
	"github.com/gin-gonic/gin"
)

type DemoService struct {
	RabbitMQService *rabbitMQ.RabbitMQService
}

func NewDemoService(rabbitMQService *rabbitMQ.RabbitMQService) *DemoService {
	return &DemoService{RabbitMQService: rabbitMQService}
}

func (s *DemoService) SendMsgWithRabbitMQ(c *gin.Context) (err error) {
	// mock a msg
	var msg = struct {
		id   int
		data string
	}{
		id:   1,
		data: "hello world",
	}

	body, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}

	//
	_ = s.RabbitMQService.Publish(mq.DemoTopic.Exchange, mq.DemoTopic.Route, body)

	return nil
}
