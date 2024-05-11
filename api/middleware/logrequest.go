package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// LogHandler log the request
func LogHandler(context *gin.Context) {
	start := time.Now()
	var requestID string

	if context.Request.Header.Get("X-Request-ID") != "" {
		requestID = context.Request.Header.Get("X-Request-ID")
	} else {
		requestID = uuid.New().String()
	}

	fields := log.Fields{
		"remote_address": context.ClientIP(),
		"user_agent":     context.Request.UserAgent(),
		"request_id":     requestID,
	}

	l := log.WithFields(fields)
	l.Infof("Request received for %s %s", context.Request.Method, context.Request.RequestURI)

	context.Next()

	cp := context.Copy()
	fields = log.Fields{
		"status":     cp.Writer.Status(),
		"size":       cp.Writer.Size(),
		"time_taken": time.Since(start),
	}

	l.WithFields(fields).Infof("Request completed for %s %s", context.Request.Method, context.Request.RequestURI)
}
