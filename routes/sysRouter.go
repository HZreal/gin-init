package routes

import (
	"github.com/gin-gonic/gin"
)

func AddSysRouter(r *gin.RouterGroup) {
	sysGroup := r.Group("sys")
	{
		sysGroup.POST("login", AppController.SysController.Login)
		sysGroup.POST("logout", AppController.SysController.Logout)
		sysGroup.POST("config")
		sysGroup.POST("config/set")
	}

}
