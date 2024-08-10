package socketio

/**
 * @Author elastic·H
 * @Date 2024-08-10
 * @File: socketio.go
 * @Description:
 */

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"log"
)

// CreateSocketIOServer 创建并配置 Socket.IO 服务器
func CreateSocketIOServer() *socketio.Server {
	// 创建一个新的 Socket.IO 服务器
	server := socketio.NewServer(nil)

	// 处理连接事件
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		s.Emit("reply", "hello "+s.ID())
		return nil
	})

	// 处理断开连接事件
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("disconnected:", s.ID(), "reason:", reason)
	})

	// 处理自定义事件
	server.OnEvent("/", "message", func(s socketio.Conn, msg string) {
		fmt.Println("received message:", msg)
		s.Emit("reply", "received: "+msg)
	})

	// 启动服务器
	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("SocketIO listen error: %s\n", err)
		}
	}()

	// 返回配置好的 Socket.IO 服务器
	return server
}
