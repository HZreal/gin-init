package core

/**
 * @Author elastic·H
 * @Date 2024-08-08
 * @File: init.go
 * @Description:
 */

import (
	"gin-init/config"
	"gin-init/core/consumer/rabbitMQ"

	// "gin-init/core/cron"
	"gin-init/core/initialize"
	gRPCServer "gin-init/core/rpc/server"
	"gin-init/core/server"
)

func Start() {
	// 加载配置
	config.Load()

	// 初始化
	initialize.Initialize()

	// 启动 mq 消费者
	rabbitMQ.StartConsumer()
	// kafka.StartConsumer()

	// 启动定时任务
	// go cron.StartCron()

	//
	gRPCServer.StartGPRC()

	//
	server.StartGinServer()
}
