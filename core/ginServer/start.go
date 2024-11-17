package ginServer

/**
 * @Author elastic·H
 * @Date 2023-06-11
 * @File: start.go
 * @Description:
 */

import (
	"fmt"
	"gin-init/config"
	"gin-init/core/socketio"
	"gin-init/middleware"
	"gin-init/routes"
	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

var r *gin.Engine

var routerGroup *gin.RouterGroup

func StartGinServer() {
	//
	gin.SetMode(config.Conf.Gin.Mode)

	// gin.Default() 使用默认的 Logger 和 Recovery
	// r = gin.Default()
	// gin 默认日志
	// r.Use(gin.Logger())
	// gin 默认异常处理 Recovery
	// r.Use(gin.Recovery())
	r = gin.New()

	// 将 Zap 日志作为 Gin 的日志中间件
	// 初始化 Logger
	if err := InitLogger(); err != nil {
		log.Printf("Failed to initialize logger: %v\n", err)
		return
	}
	r.Use(ginzap.Ginzap(Logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(Logger, true))

	// 异常拦截处理，代替默认的 Recovery
	r.Use(middleware.ExceptionInterceptorMiddleware())

	// 注册 CORS 中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                       // 允许的域名
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},            // 允许的HTTP方法
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // 允许的请求头
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,           // 允许携带认证信息
		MaxAge:           12 * time.Hour, // 预检请求的缓存时间
	}))

	// 创建并配置 Socket.IO 服务器
	socketIOServer := socketio.CreateSocketIOServer()
	// 将 Socket.IO 服务器集成到 Gin 路由中
	r.GET("/socket.io/*any", gin.WrapH(socketIOServer))
	r.POST("/socket.io/*any", gin.WrapH(socketIOServer))

	// 静态文件（页面等）
	r.Static("/public", "./public")

	// router 组定义
	routerGroup = r.Group("api")
	// 注册路由
	routes.RegisterRoutes(routerGroup)

	// 启动
	err := r.Run(config.Conf.Gin.GetAddr())
	if err != nil {
		fmt.Println("[Error] r.Run " + err.Error())
		return
	}
}
