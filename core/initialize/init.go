package initialize

import (
	"gin-init/core/initialize/database"
	"gin-init/core/initialize/logServer"
	"gin-init/core/initialize/mq/rabbitMQ"
)

func Initialize() {
	//
	database.InitMysql()

	//
	database.InitRedis()

	//
	rabbitMQ.InitRabbitMQ()

	//
	// database.InitPostgresql()

	// log server
	logServer.InitLogServer()
}
