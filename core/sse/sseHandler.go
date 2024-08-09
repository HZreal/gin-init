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
