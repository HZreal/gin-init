package routes

/**
 * @Author elastic·H
 * @Date 2024-09-29
 * @File: routes.go
 * @Description:
 */

import (
	"gin-init/core/wire"
	"github.com/gin-gonic/gin"
)

var (
	AppController *wire.AppControllers
)

var routerRegistrars = []RouterRegistrar{
	&DemoRouter{},
	&SysRouter{},
	&UserRouter{},
	// 可以继续添加更多的路由组
}

// RouterRegistrar 路由组注册接口
type RouterRegistrar interface {
	RegisterRoutes(r *gin.RouterGroup)
}

func RegisterRoutes(routerGroup *gin.RouterGroup) {
	//
	var err error
	AppController, err = wire.InitializeApp()
	if err != nil {
		panic(err)
	}

	for _, registrar := range routerRegistrars {
		registrar.RegisterRoutes(routerGroup)
	}
}
