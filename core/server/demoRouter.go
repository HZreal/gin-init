package server

import (
	"gin-init/core/sse"
	"gin-init/core/wire"
	"gin-init/core/ws"
)

func addDemoRouter() {
	// demoController := controller.DemoController{}
	appController, _ := wire.InitializeApp()

	sysGroup := apiGroup.Group("demo")
	{
		sysGroup.GET("sse", sse.SseHandler)
		sysGroup.GET("ws", ws.WebsocketHandler)
		sysGroup.GET("sendMQ", appController.DemoController.SendMsgWithRabbitMQ)
	}

}
