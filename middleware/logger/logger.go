package logger

import (
	"io"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

const Time_FMT = "2006-01-02 15:04:05"
const LOG_APP_FILE = "app.log"
const LOG_SERVER_FILE = "server.log"

func initLogrusConf(p string) {
	logFile := &lumberjack.Logger{
		Filename:   path.Join(p, LOG_SERVER_FILE), // 日志文件名
		MaxSize:    50,                            // 单个日志文件最大尺寸（单位：MB）
		MaxBackups: 10,                            // 保留旧日志文件的最大个数
		MaxAge:     180,                           // 保留旧日志文件的最大天数
		Compress:   true,                          // 是否启用压缩
		LocalTime:  true,                          // 使用本地时间（默认为UTC时间）
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(io.MultiWriter(os.Stdout, logFile))
}

func LoggerToFile(p string) gin.HandlerFunc {
	initLogrusConf(p)

	logger := logrus.New()
	logFile := &lumberjack.Logger{
		Filename:   path.Join(p, LOG_APP_FILE), // 日志文件名
		MaxSize:    50,                         // 单个日志文件最大尺寸（单位：MB）
		MaxBackups: 10,                         // 保留旧日志文件的最大个数
		MaxAge:     180,                        // 保留旧日志文件的最大天数
		Compress:   true,                       // 是否启用压缩
		LocalTime:  true,                       // 使用本地时间（默认为UTC时间）
	}

	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(io.MultiWriter(os.Stdout, logFile))

	return func(c *gin.Context) {
		startTime := time.Now()               // start time
		c.Next()                              // Handling the Request
		endTime := time.Now()                 // end time
		latencyTime := endTime.Sub(startTime) // execution time
		reqMethod := c.Request.Method         // request method
		reqUri := c.Request.RequestURI        // required parameter
		statusCode := c.Writer.Status()       // status code
		clientIP := c.ClientIP()              // IP
		logger.Infof("[GIN] %13v | %15s | %7s | %3d | %s ",
			latencyTime,
			clientIP,
			reqMethod,
			statusCode,
			reqUri,
		)
	}
}
