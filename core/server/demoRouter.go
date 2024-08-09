package server

import (
	"gin-init/core/sse"
)

func addDemoRouter() {
	sysGroup := apiGroup.Group("demo")
	{
		sysGroup.GET("sse", sse.SseHandler)
	}

}
