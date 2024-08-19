// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"gin-init/controller"
	"gin-init/model/entity"
	"gin-init/service"
	"github.com/google/wire"
)

// Injectors from wire.go:

// InitializeApp 初始化整个应用的控制器
func InitializeApp() (*AppControllers, error) {
	userModel := entity.NewUserModel()
	userService := service.NewUserService(userModel)
	userController := controller.NewUserController(userService)
	sysService := service.NewSysService(userService)
	sysController := controller.NewSysController(sysService)
	appControllers := &AppControllers{
		UserController: userController,
		SysController:  sysController,
	}
	return appControllers, nil
}

// wire.go:

var UserSet = wire.NewSet(entity.NewUserModel, service.NewUserService, controller.NewUserController)

var SysSet = wire.NewSet(service.NewSysService, controller.NewSysController)

// AppSet 包含了所有模型的 ProviderSet
var AppSet = wire.NewSet(
	UserSet,
	SysSet,
)

type AppControllers struct {
	UserController *controller.UserController
	SysController  *controller.SysController
}
