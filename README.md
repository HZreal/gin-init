# gin-init 项目初始化模板

## 简介

基于 gin 的二次开发，集成了一些基本功能、组件的模板

助你快速启动一个初始化完备的 gin 应用



本模板基于 gin@1.10，详情请查看 https://github.com/gin-gonic/gin



## 本模板已支持的基本能力、组件

- 基础路由、基础查询、错误码定义与通用响应
- 跨域 CORS
- 参数验证
- 异常拦截器
- JWT 认证
- Gorm
- Mysql、Postgresql 数据库驱动支持
- Redis
- RabbitMQ 消息队列
- Cron 定时
- gRPC 支持 RPC 调用
- Wire 依赖注入
- Websocket、socketIO、SSE 支持
- Viper 多环境配置
- Zap 日志
- Samberlo、Carbon 等效率工具
- 链路追逐



## 快速使用

##### 安装依赖

```bash
$ go mod tidy
```

##### 检查支持的能力、组件及配置

1. 检查需要启用的组件，查看 core/init.ts 文件
2. 检查需要启用组件的调整配置，查看 config/xxx.yaml 文件

##### 启动运行

```
$ go run main.go
$ open http://localhost:40020
```



## 启用默认禁用的组件能力说明

##### 启用 Cron、rabbitMQ、gRPC 等

core.Start() 方法中启动



##### 启用 sse 协议

看一下函数 SseHandler 的示例即可



## PS

本模板部分暂未集成的能力和组件，后续添补
