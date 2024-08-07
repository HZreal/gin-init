package server

import (
	"fmt"
	"gin-init/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

var r *gin.Engine

var apiGroup *gin.RouterGroup

func init() {
	//
	gin.SetMode(config.Conf.Gin.Mode)

	//
	r = gin.Default()

	//
	// r.Use(gin.Logger())

	//
	// r.Use(gin.Recovery())

	// cors
	// 全局注册 CORS 中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                       // 允许的域名
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},            // 允许的HTTP方法
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // 允许的请求头
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,           // 允许携带认证信息
		MaxAge:           12 * time.Hour, // 预检请求的缓存时间
	}))

	// routers definition
	apiGroup = r.Group("api")

	//
	registerRoutes()
}

func registerRoutes() {
	addSysRouter()
	addDemoRouter()
	addUserRouter()
}

func StartGinServer() {
	err := r.Run(config.Conf.Gin.GetAddr())
	if err != nil {
		fmt.Println("[Error] r.Run " + err.Error())
		return
	}
}
