package initialize

import (
	"gin-init/core/initialize/database"
	"gin-init/core/initialize/logServer"
	"gin-init/core/initialize/mq"
)

func Initialize() {
	//
	database.InitMysql()

	//
	database.InitRedis()

	//
	mq.InitRabbitMQ()
	mq.InitKafka()

	//
	// database.InitPostgresql()

	// log server
	logServer.InitLogServer()
}
