package ginServer

/**
 * @Author elastic·H
 * @Date 2024-09-08
 * @File: zapLogger.go
 * @Description:
 */

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin-init/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/http"
	"os"
	"time"
)

var Logger *zap.Logger

// 基础日志器
func init2() {
	// 开发环境下使用 zap.NewDevelopment()，生产环境下使用 zap.NewProduction()
	Logger, _ = zap.NewDevelopment()
	// Logger, _ = zap.NewProduction()
}

// 日志输出到文件
func init3() {
	logFile, err := os.OpenFile("logs/gin_app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	// 创建自定义的 zapcore，输出到文件
	writeSyncer := zapcore.AddSync(logFile)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // 输出 JSON 格式的日志
		writeSyncer,                           // 输出到日志文件
		zapcore.InfoLevel,                     // 日志级别
	)

	Logger = zap.New(core, zap.AddCaller())
}

// InitLogger 根据配置初始化 zap.Logger
func InitLogger() error {
	logConfig := config.Conf.Logger

	// 编码器配置
	encoderConfig := zapcore.EncoderConfig{
		// MessageKey:    "", // 请求 path
		LevelKey: "level",
		// TimeKey:       "",
		// NameKey:       "",
		CallerKey:     "caller",
		FunctionKey:   "",
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		EncodeTime:    zapcore.ISO8601TimeEncoder,
		EncodeCaller:  zapcore.ShortCallerEncoder,
		StacktraceKey: "stacktrace",
	}

	encoder := zapcore.NewJSONEncoder(encoderConfig)

	// 日志核心集合
	var cores []zapcore.Core

	// 控制台日志
	if logConfig.Console.Enable {
		consoleLevel := getLogLevel(logConfig.Console.Level)
		consoleCore := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), consoleLevel)
		cores = append(cores, consoleCore)
	}

	// 文件日志
	if logConfig.File.Enable {
		fileLevel := getLogLevel(logConfig.File.Level)
		fileWriter, err := getFileWriter(logConfig.File.Path, logConfig.File.MaxSize, logConfig.File.MaxBackupCount)
		if err != nil {
			return fmt.Errorf("failed to create file writer: %v", err)
		}
		fileCore := zapcore.NewCore(encoder, fileWriter, fileLevel)
		cores = append(cores, fileCore)
	}

	// 外部日志收集器服务
	if logConfig.External.Enable {
		externalLevel := getLogLevel(logConfig.External.Level)
		externalCore, err := getExternalCore(logConfig.External.Service, logConfig.External.Url, logConfig.External.Index, encoder, externalLevel)
		if err != nil {
			return fmt.Errorf("failed to create external core: %v", err)
		}
		cores = append(cores, externalCore)
	}

	// 创建日志器
	core := zapcore.NewTee(cores...)
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return nil
}

// getLogLevel 将字符串级别转换为 zapcore.Level
func getLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

// getFileWriter 创建文件日志写入器
func getFileWriter(filepath string, maxSize int, maxBackupCount int) (zapcore.WriteSyncer, error) {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   filepath,
		MaxSize:    maxSize / (1024 * 1024), // 转换为 MB
		MaxBackups: maxBackupCount,
		Compress:   true, // 启用压缩
	}
	return zapcore.AddSync(lumberjackLogger), nil
}

// ElkWriter 定义发送到 ELK 的日志写入器
// TODO 待实现测试
type ElkWriter struct {
	URL   string // ElasticSearch 的 URL
	Index string // 日志索引名
}

// Write 实现 zapcore.WriteSyncer 接口，将日志发送到 ELK
func (e *ElkWriter) Write(p []byte) (n int, err error) {
	// 构造 ElasticSearch 文档格式
	logEntry := map[string]interface{}{
		"@timestamp": time.Now().UTC().Format(time.RFC3339), // 时间戳
		"message":    string(p),                             // 原始日志内容
	}

	// 转换为 JSON
	body, err := json.Marshal(logEntry)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal log entry: %v", err)
	}

	// 发送 HTTP 请求到 ElasticSearch
	url := fmt.Sprintf("%s/%s/_doc", e.URL, e.Index)
	resp, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return 0, fmt.Errorf("failed to send log to ELK: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		// 返回写入字节数
		return len(p), nil
	}

	return 0, fmt.Errorf("elk returned non-success status: %s", resp.Status)
}

// getExternalCore 配置外部日志服务
func getExternalCore(service, url, index string, encoder zapcore.Encoder, level zapcore.Level) (zapcore.Core, error) {
	if service == "elk" {
		writer := zapcore.AddSync(&ElkWriter{
			URL:   url,
			Index: index,
		})
		return zapcore.NewCore(encoder, writer, level), nil
	}

	return nil, fmt.Errorf("unsupported external service: %s", service)
}
