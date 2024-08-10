package ws

/**
 * @Author elastic·H
 * @Date 2024-08-10
 * @File: websocket.go
 * @Description:
 */

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

// 升级 HTTP 连接为 WebSocket 连接的 Upgrader
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有连接
	},
}

func WebsocketHandler(c *gin.Context) {
	// 将 HTTP 连接升级为 WebSocket 连接
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to upgrade to WebSocket"})
		return
	}
	defer conn.Close()

	// 简单的消息循环
	for {
		// 读取消息
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Read error:", err)
			break
		}

		fmt.Printf("Received: %s\n", message)

		// 回复消息
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			fmt.Println("Write error:", err)
			break
		}
	}
}
