package middleware

import (
	"time"
	"go.uber.org/zap"
	"github.com/gin-gonic/gin"
	"github.com/Sina-Mahmoodmoradi/ecommerce-backend/internal/pkg/logger"
)

func GinLogger(l *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		reqID, _ := c.Get("RequestID")

		l.Info("HTTP Request",
			zap.Int("status", status),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Duration("latency", latency),
			zap.String("request_id", reqID.(string)),
		)
	}
}
