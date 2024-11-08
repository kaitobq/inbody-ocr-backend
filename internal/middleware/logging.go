package middleware

import (
	"inbody-ocr-backend/internal/domain/xcontext"
	"inbody-ocr-backend/internal/infra/logging"

	"github.com/gin-gonic/gin"
)

func LoggingRequestTraceID() gin.HandlerFunc {
	return func(c *gin.Context) {
		xcontext.WithLoggingRequestTraceID(c, logging.GetRequestTraceID(c))

		c.Next()
	}
}
