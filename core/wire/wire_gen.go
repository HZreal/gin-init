// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"gin-init/controller"
	"gin-init/model/entity"
	"gin-init/mq/rabbitMQ"
	"gin-init/service"
	"github.com/google/wire"
)

// Injectors from wire.go:

// InitializeApp 初始化整个应用的控制器
func InitializeApp() (*AppControllers, error) {
	rabbitMQService := rabbitMQ.NewRabbitMQService()
	demoService := service.NewDemoService(rabbitMQService)
	demoController := controller.NewDemoController(demoService)
	userModel := entity.NewUserModel()
	userService := service.NewUserService(userModel)
	sysService := service.NewSysService(userService)
	sysController := controller.NewSysController(sysService)
	userController := controller.NewUserController(userService)
	appControllers := &AppControllers{
		DemoController: demoController,
		SysController:  sysController,
		UserController: userController,
	}
	return appControllers, nil
}

// wire.go:

var DemoSet = wire.NewSet(controller.NewDemoController, service.NewDemoService, rabbitMQ.NewRabbitMQService)

var SysSet = wire.NewSet(controller.NewSysController, service.NewSysService)

var UserSet = wire.NewSet(controller.NewUserController, service.NewUserService, entity.NewUserModel)

// AppSet 包含了所有模型的 ProviderSet
var AppSet = wire.NewSet(
	DemoSet,
	SysSet,
	UserSet,
)

type AppControllers struct {
	DemoController *controller.DemoController
	SysController  *controller.SysController
	UserController *controller.UserController
}
