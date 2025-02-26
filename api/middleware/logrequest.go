package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// LogHandler logs the request and response completion events
func LogHandler(context *gin.Context) {
	start := time.Now()
	var requestID string

	// Check if the request ID is already set in the header
	if context.Request.Header.Get("X-Request-ID") != "" {
		requestID = context.Request.Header.Get("X-Request-ID")
	} else {
		requestID = uuid.New().String()
	}

	// Set fields for the log entry
	fields := log.Fields{
		"remote_address": context.ClientIP(),
		"user_agent":     context.Request.UserAgent(),
		"request_id":     requestID,
	}

	// Create a new log entry with the fields
	l := log.WithFields(fields)
	// Log the request received event
	l.Infof("Request received for %s %s", context.Request.Method, context.Request.RequestURI)

	context.Next()

	cp := context.Copy()
	fields = log.Fields{
		"status":     cp.Writer.Status(),
		"size":       cp.Writer.Size(),
		"time_taken": time.Since(start),
	}

	// Log the request completed event
	l.WithFields(fields).Infof("Request completed for %s %s", context.Request.Method, context.Request.RequestURI)
}
