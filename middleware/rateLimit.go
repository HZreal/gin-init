package middleware

/**
 * @Author elastic·H
 * @Date 2024-08-27
 * @File: rateLimit.go
 * @Description:
 */

import (
	"gin-init/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// RateLimitMiddleware 接口频率限制
func RateLimitMiddleware(maxRequests int, resetTime time.Duration) gin.HandlerFunc {
	var requests = 0
	var lastReset = time.Now()

	return func(c *gin.Context) {
		// 检查是否需要重置请求计数
		if time.Since(lastReset) > resetTime {
			requests = 0
			lastReset = time.Now()
		}

		if requests >= maxRequests {
			// 如果超出请求限制，返回 429 错误
			c.JSON(http.StatusTooManyRequests, common.Failed(common.TooManyRequestsError))
			c.Abort()
		}

		requests++

		//
		c.Next()
	}
}