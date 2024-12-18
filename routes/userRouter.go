package routes

import (
	"gin-init/controller"
	"gin-init/middleware"
	"github.com/gin-gonic/gin"
	"time"
)

type UserRouter struct{}

func (d *UserRouter) RegisterRoutes(r *gin.RouterGroup) {
	userGroup := r.Group("user")
	{
		userGroup.POST("info/all", AppController.UserController.GetAllUser)
		userGroup.GET("info/detail/2", controller.NewUserController2().GetByID)
		userGroup.GET("info/detail", middleware.RateLimitMiddleware(5, time.Minute), AppController.UserController.GetUserDetail)
		// userGroup.POST("info/list", middleware.JWTMiddleware(), userController.GetUserList)
		userGroup.POST("info/list", middleware.JWTMiddleware(), AppController.UserController.GetUserList)
		userGroup.POST("info/create/2", controller.NewUserController2().Create)
		userGroup.POST("info/create", AppController.UserController.CreateUser)
		userGroup.POST("info/update", AppController.UserController.UpdateUser)
		userGroup.POST("info/update/passwd", AppController.UserController.UpdateUserPassword)
		userGroup.POST("info/reset/passwd", AppController.UserController.ResetUserPassword)
		userGroup.POST("info/delete", AppController.UserController.DeleteUser)
	}
}
