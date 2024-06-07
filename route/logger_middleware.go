package route

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		startTime := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		// Process request
		c.Next()
		param := gin.LogFormatterParams{
			Request: c.Request,
			Keys:    c.Keys,
		}
		// Stop timer
		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(startTime)

		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		param.BodySize = c.Writer.Size()

		if raw != "" {
			path = path + "?" + raw
		}

		param.Path = path
		msg := defaultLogFormatter(param)
		if len(c.Errors) > 0 { // 如果存在内部错误
			slog.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if param.StatusCode >= 500 {
			slog.Error(msg)
		} else if param.StatusCode >= 400 {
			slog.Warn(msg)
		}
		// else {
		// 	// log.Info(msg)
		// }
	}
}

// defaultLogFormatter is the default log format function Logger middleware uses.
var defaultLogFormatter = func(param gin.LogFormatterParams) string {
	if param.Latency > time.Minute {
		param.Latency = param.Latency - param.Latency%time.Second
	}
	return fmt.Sprintf(" |%3d| %13v | %13s | %6v |%-7s | %#v | %s",
		param.StatusCode,
		param.Latency,
		param.ClientIP,
		param.BodySize,
		param.Method,
		param.Path,
		param.ErrorMessage,
	)
}
