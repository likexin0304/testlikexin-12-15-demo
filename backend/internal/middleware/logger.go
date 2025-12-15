package middleware

import (
	"time"

	"{{.ProjectName}}/pkg/logger"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		latency := time.Since(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		if raw != "" {
			path = path + "?" + raw
		}

		logger.Info("%s | %3d | %13v | %15s | %s",
			method,
			statusCode,
			latency,
			clientIP,
			path,
		)
	}
}

