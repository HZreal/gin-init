package middleware

/**
 * @Author elastic·H
 * @Date 2024-08-27
 * @File: exceptionInterceptor.go
 * @Description:
 */

import (
	"gin-init/common"
	"github.com/gin-gonic/gin"
)

// ExceptionInterceptorMiddleware 全局异常拦截器
func ExceptionInterceptorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		defer func() {
			if err := recover(); err != nil {
				// 检查是否为 ErrorCode 类型
				if errCode, ok := err.(common.ErrorCode); ok {
					// 返回自定义错误码和错误信息
					common.Failed(c, errCode)
				} else if msg, ok2 := err.(string); ok2 {
					// 若 panic 一个字符串
					common.FailedWithMsg(c, msg)
				} else {
					// 处理其他未知的 panic
					common.Failed(c, common.UnKnownError)
				}
				c.Abort()
			}
		}()

		c.Next()

	}
}

func ExceptionMiddleware(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			// 简单返回友好提示，具体可自定义发生错误后处理逻辑
			common.Failed(c, common.UnKnownError)
			c.Abort()
		}
	}()

	c.Next()
}
