package xcontext

import "github.com/gin-gonic/gin"

var (
	loggingRequestTraceIDKey = "LOGGING:REQUEST_TRACE_ID"
)

func WithLoggingRequestTraceID(c *gin.Context, traceID string) {
	c.Set(loggingRequestTraceIDKey, traceID)
}

func LoggingRequestTraceID(c *gin.Context) string {
	if v, ok := c.Get(loggingRequestTraceIDKey); ok {
		return v.(string)
	}
	return ""
}
