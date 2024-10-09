//go:build wireinject
// +build wireinject

package wire

import (
	"gin-init/controller"
	"gin-init/model/entity"
	"gin-init/mq/rabbitMQ"
	"gin-init/service"
	"gin-init/service/common"
	"github.com/google/wire"
)

// 在本文件所在目录下执行 wire 命令

var DemoSet = wire.NewSet(
	controller.NewDemoController,
	service.NewDemoService,
	rabbitMQ.NewRabbitMQService,
)

var SysSet = wire.NewSet(
	controller.NewSysController,
	service.NewSysService,
)

var UserSet = wire.NewSet(
	controller.NewUserController,
	service.NewUserService,
	common.NewRedisService,
	entity.NewUserModel,
)

// AppSet 包含了所有模型的 ProviderSet
var AppSet = wire.NewSet(
	DemoSet,
	SysSet,
	UserSet,
	// 可以在这里继续添加其他模块的 ProviderSet
)

type AppControllers struct {
	DemoController *controller.DemoController
	SysController  *controller.SysController
	UserController *controller.UserController
}

// InitializeApp 初始化整个应用的控制器
func InitializeApp() (*AppControllers, error) {
	wire.Build(AppSet, wire.Struct(new(AppControllers), "*"))
	return &AppControllers{}, nil
}
