package initialize

import (
	"gin-init/core/initialize/database"
	"gin-init/core/initialize/logServer"
)

func Initialize() {
	//
	database.InitMysql()

	//
	database.InitRedis()

	//
	// database.InitPostgresql()

	// log server
	logServer.InitLogServer()
}
