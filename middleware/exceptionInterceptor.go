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
	"net/http"
)

// ExceptionInterceptorMiddleware 全局异常拦截器
func ExceptionInterceptorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		defer func() {
			if err := recover(); err != nil {
				// 检查是否为 ErrorCode 类型
				if errCode, ok := err.(common.ErrorCode); ok {
					// 返回自定义错误码和错误信息
					c.JSON(http.StatusInternalServerError, common.Failed(errCode))
				} else if msg, ok2 := err.(string); ok2 {
					// 若 panic 一个字符串
					c.JSON(http.StatusInternalServerError, common.FailedWithMsg(msg))
				} else {
					// 处理其他未知的 panic
					c.JSON(http.StatusInternalServerError, common.Failed(common.UnKnownError))
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
			c.JSON(http.StatusInternalServerError, common.Failed(common.UnKnownError))
			c.Abort()
		}
	}()

	c.Next()
}
