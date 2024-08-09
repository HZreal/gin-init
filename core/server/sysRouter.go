package server

import "gin-init/controller"

func addSysRouter() {
	sysController := controller.SysController{}
	sysGroup := apiGroup.Group("sys")
	{
		sysGroup.POST("login", sysController.Login)
		sysGroup.POST("logout", sysController.Logout)
		sysGroup.POST("config")
		sysGroup.POST("config/set")
	}

}
