package server

import (
	"gin-init/core/wire"
	"gin-init/middleware"
	"time"
)

func addUserRouter() {
	// userController := controller.UserController{}
	appController, _ := wire.InitializeApp()

	userGroup := apiGroup.Group("user")
	{
		userGroup.GET("info/all")
		// userGroup.GET("info/detail", userController.GetUserDetail)
		userGroup.GET("info/detail", middleware.RateLimitMiddleware(5, time.Minute), appController.UserController.GetUserDetail)
		// userGroup.POST("info/list", middleware.JWTMiddleware(), userController.GetUserList)
		userGroup.POST("info/list", middleware.JWTMiddleware(), appController.UserController.GetUserList)
		// userGroup.POST("info/create", userController.CreateUser)
		userGroup.POST("info/create", appController.UserController.CreateUser)
		userGroup.POST("info/update")
		userGroup.POST("info/reset/passwd")
		userGroup.POST("info/delete")
	}
}
