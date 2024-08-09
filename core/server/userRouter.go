package server

import (
	"gin-init/controller"
	"gin-init/middleware"
)

func addUserRouter() {
	userController := controller.UserController{}

	userGroup := apiGroup.Group("user")
	{
		userGroup.GET("info/all")
		userGroup.GET("info/detail", userController.GetUserDetail)
		userGroup.POST("info/list", middleware.JWTMiddleware(), userController.GetUserList)
		userGroup.POST("info/create", userController.CreateUser)
		userGroup.POST("info/update")
		userGroup.POST("info/reset/passwd")
		userGroup.POST("info/delete")
	}
}
