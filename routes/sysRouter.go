package routes

import (
	"github.com/gin-gonic/gin"
)

type SysRouter struct{}

func (d SysRouter) RegisterRoutes(r *gin.RouterGroup) {
	sysGroup := r.Group("sys")
	{
		sysGroup.POST("login", AppController.SysController.Login)
		sysGroup.POST("logout", AppController.SysController.Logout)
		sysGroup.POST("config")
		sysGroup.POST("config/set")
	}
}

func AddSysRouter(r *gin.RouterGroup) {
	sysGroup := r.Group("sys")
	{
		sysGroup.POST("login", AppController.SysController.Login)
		sysGroup.POST("logout", AppController.SysController.Logout)
		sysGroup.POST("config")
		sysGroup.POST("config/set")
	}
}
