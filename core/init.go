package core

/**
 * @Author elastic·H
 * @Date 2024-08-08
 * @File: init.go
 * @Description:
 */

import (
	gRPCServer "gin-init/core/rpc/server"
	"gin-init/core/server"
	// "gin-init/job"
	"gin-init/mq"
)

func Start() {
	//
	mq.Start()

	//
	// go job.StartCron()

	//
	gRPCServer.StartGPRC()

	//
	server.StartGinServer()
}
