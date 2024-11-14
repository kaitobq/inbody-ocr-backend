package logging

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func GetRequestTraceID(c *gin.Context) string {
	header := c.Request.Header.Get("X-Cloud-Trace-Context")
	if header == "" {
		return ""
	}
	parts := strings.Split(header, "/")
	if len(parts) == 0 {
		return ""
	}
	return parts[0]
}
