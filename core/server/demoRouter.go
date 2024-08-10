package server

import (
	"gin-init/core/sse"
	"gin-init/core/ws"
)

func addDemoRouter() {
	sysGroup := apiGroup.Group("demo")
	{
		sysGroup.GET("sse", sse.SseHandler)
		sysGroup.GET("ws", ws.WebsocketHandler)

	}

}
