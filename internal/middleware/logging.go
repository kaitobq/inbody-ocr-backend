package middleware

import (
	"inbody-ocr-backend/internal/domain/xcontext"
	"inbody-ocr-backend/internal/infra/logging"

	"github.com/gin-gonic/gin"
)

type Logging struct{}

func NewLogging() *Logging {
	return &Logging{}
}

func (l *Logging) LoggingRequestTraceID() gin.HandlerFunc {
	return func(c *gin.Context) {
		xcontext.WithLoggingRequestTraceID(c, logging.GetRequestTraceID(c))

		c.Next()
	}
}
