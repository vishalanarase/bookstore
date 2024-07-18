package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/vishalanarase/bookstore/pkg/monitoring"
)

// MetricsMiddleware returns a Gin middleware function that increments request metrics
func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method

		timer := prometheus.NewTimer(monitoring.RequestDuration.WithLabelValues(path, method))
		defer timer.ObserveDuration()

		monitoring.RequestCounter.WithLabelValues(path, method).Inc()

		c.Next()
	}
}
