package routes

/**
 * @Author elastic·H
 * @Date 2024-09-29
 * @File: init.go
 * @Description:
 */

import (
	"gin-init/core/wire"
)

var AppController *wire.AppControllers

func init() {
	var err error
	AppController, err = wire.InitializeApp()
	if err != nil {
		panic(err)
	}
}
