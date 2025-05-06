package sse

/**
 * @Author elastic·H
 * @Date 2024-08-09
 * @File: sseHandler.go
 * @Description:
 */

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

func SseHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			message := fmt.Sprintf("data: The time is %s\n\n", t.Format(time.RFC3339))
			_, err := c.Writer.Write([]byte(message))
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			c.Writer.Flush()
		case <-c.Request.Context().Done():
			return
		}
	}
}

// --------------------------------------------------------------------------------------------------------
// 封装
// 1. 创建一个 Hub 结构体，用于管理所有的客户端连接；在 Hub 中维护一个 map，用于存储所有的客户端连接
// 2. 提供一个方法，用于向所有客户端广播消息

type SseHub struct {
	clients map[chan string]bool
	lock    sync.RWMutex
}

var Hub = NewSseHub() // 单例

func NewSseHub() *SseHub {
	return &SseHub{
		clients: make(map[chan string]bool),
	}
}

func (h *SseHub) Serve(c *gin.Context) {
	w := c.Writer
	r := c.Request

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	clientChan := make(chan string, 10)
	h.addClient(clientChan)
	defer h.removeClient(clientChan)

	notify := r.Context().Done()
	flusher, _ := w.(http.Flusher)

	for {
		select {
		case msg := <-clientChan:
			fmt.Fprintf(w, "data: %s\n\n", msg)
			flusher.Flush()
		case <-notify:
			return
		}
	}
}

func (h *SseHub) addClient(ch chan string) {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.clients[ch] = true
}

func (h *SseHub) removeClient(ch chan string) {
	h.lock.Lock()
	defer h.lock.Unlock()
	delete(h.clients, ch)
	close(ch)
}

func (h *SseHub) Broadcast(msg string) {
	h.lock.RLock()
	defer h.lock.RUnlock()
	for ch := range h.clients {
		select {
		case ch <- msg:
		default:
		}
	}
}

// SimulateEvents 模拟周期性事件
func (h *SseHub) SimulateEvents() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		message := fmt.Sprintf("Server time: %s", time.Now().Format(time.RFC3339))
		// 广播给所有客户端
		for clientChan := range h.clients {
			select {
			case clientChan <- message:
			default:
				// 跳过阻塞的 channel
			}
		}
	}
}
