package mq

/**
 * @Author elastic·H
 * @Date 2024-08-08
 * @File: init.go
 * @Description:
 */

import (
	"gin-init/mq/kafka"
	"gin-init/mq/rabbitMQ"
)

func Start() {
	//
	rabbitMQ.Start()

	//
	kafka.Start()
}
