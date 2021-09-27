package middleware

import (
	"io"
	"os"
	"time"

	"github.com/91go/rss2/utils/log"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

const (
	LogMaxSize    = 20
	LogMaxBackups = 5
	LogMaxAge     = 30
)

func Logger() gin.HandlerFunc {
	// 设置日志格式为json格式
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// 设置日志级别为warn以上
	logrus.SetLevel(logrus.WarnLevel)
	logrus.SetReportCaller(true)

	path := "./logrus.log"

	// 设置将日志输出到标准输出（默认的输出为stderr,标准错误）
	// 日志消息输出可以是任意的io.writer类型
	logrus.SetOutput(rolling(path))

	dh, _ := log.NewDingHook(log.AssembleUrl(), nil)

	logrus.AddHook(dh)

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUrl := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求ip
		clientIP := c.ClientIP()

		// 日志格式
		logrus.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUrl,
		}).Info()
	}
}

func rolling(path string) io.Writer {
	logger := &lumberjack.Logger{
		LocalTime:  true,
		Filename:   path,          // 日志文件位置
		MaxSize:    LogMaxSize,    // 单文件最大容量,单位是MB
		MaxBackups: LogMaxBackups, // 最大保留过期文件个数
		MaxAge:     LogMaxAge,     // 保留过期文件的最大时间间隔,单位是天
		Compress:   false,         // 是否需要压缩滚动日志, 使用的 gzip 压缩
	}
	writers := []io.Writer{
		logger,
		os.Stdout,
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	return fileAndStdoutWriter
}
