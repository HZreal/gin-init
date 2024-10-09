package routes

import (
	"gin-init/core/sse"
	"gin-init/core/ws"
	"github.com/gin-gonic/gin"
)

type DemoRouter struct{}

func (d *DemoRouter) RegisterRoutes(r *gin.RouterGroup) {
	sysGroup := r.Group("demo")
	{
		sysGroup.GET("sse", sse.SseHandler)
		sysGroup.GET("ws", ws.WebsocketHandler)
		sysGroup.GET("sendMQ", AppController.DemoController.SendMsgWithRabbitMQ)
	}
}
