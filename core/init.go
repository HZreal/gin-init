package core

/**
 * @Author elastic·H
 * @Date 2024-08-08
 * @File: init.go
 * @Description:
 */

import (
	"gin-init/config"
	"gin-init/core/initialize"
	gRPCServer "gin-init/core/rpc/server"
	"gin-init/core/server"
	"gin-init/mq"
)

func Start() {
	// 加载配置
	config.Load()

	// 初始化
	initialize.Initialize()

	//
	mq.Start()

	//
	// go job.StartCron()

	//
	gRPCServer.StartGPRC()

	//
	server.StartGinServer()
}
