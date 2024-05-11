package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/vishalanarase/bookstore/pkg/metrics"
)

// MetricsHandlerMiddleware returns a Gin middleware function that increments request metrics
func MetricsHandlerMiddleware(metrics *metrics.Metrics) gin.HandlerFunc {
	return func(context *gin.Context) {
		// Increment request counter using metrics object
		metrics.IncrementRequest(context.Request.Method)

		// Call the next handler
		context.Next()
	}
}
