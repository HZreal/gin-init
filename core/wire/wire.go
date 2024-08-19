//go:build wireinject
// +build wireinject

package wire

import (
	"gin-init/controller"
	"gin-init/model/entity"
	"gin-init/service"
	"github.com/google/wire"
)

var UserSet = wire.NewSet(
	entity.NewUserModel,
	service.NewUserService,
	controller.NewUserController,
)

var SysSet = wire.NewSet(
	service.NewSysService,
	controller.NewSysController,
)

// AppSet 包含了所有模型的 ProviderSet
var AppSet = wire.NewSet(
	UserSet,
	SysSet,
	// 可以在这里继续添加其他模块的 ProviderSet
)

type AppControllers struct {
	UserController *controller.UserController
	SysController  *controller.SysController
}

// InitializeApp 初始化整个应用的控制器
func InitializeApp() (*AppControllers, error) {
	wire.Build(AppSet, wire.Struct(new(AppControllers), "*"))
	return &AppControllers{}, nil
}
