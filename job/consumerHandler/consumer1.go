package consumerHandler

/**
 * @Author huang
 * @Date 2024-07-12
 * @File: consumer1.go
 * @Description:
 */

import (
	"gin-init/common/constant"
	"log"
)

func HandleMessage1(msg []byte) {
	log.Printf("Received a message from queue1: %s", msg)
}

type AAAConsumer struct {
	Topic constant.Topic
}

func NewAAAConsumer() *AAAConsumer {
	return &AAAConsumer{
		Topic: constant.DemoTopic,
	}
}

func (c *AAAConsumer) HandleMessage(msg []byte) {
	log.Printf("Received a message from queue1: %s", msg)

	//
}
