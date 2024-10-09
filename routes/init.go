package routes

/**
 * @Author elastic·H
 * @Date 2024-09-29
 * @File: init.go
 * @Description:
 */

import (
	"gin-init/core/wire"
	"github.com/gin-gonic/gin"
)

var (
	AppController    *wire.AppControllers
	routerRegistrars []RouterRegistrar
)

func init() {
	//
	var err error
	AppController, err = wire.InitializeApp()
	if err != nil {
		panic(err)
	}

	//
	routerRegistrars = []RouterRegistrar{
		&DemoRouter{},
		&SysRouter{},
		&UserRouter{},
		// 可以继续添加更多的路由组
	}
}

// RouterRegistrar 路由组注册接口
type RouterRegistrar interface {
	RegisterRoutes(r *gin.RouterGroup)
}

// GetRouterRegistrars 获取所有注册器
func GetRouterRegistrars() []RouterRegistrar {
	return routerRegistrars
}
