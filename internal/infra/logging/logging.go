package logging

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
)

var l *slog.Logger

func Init() {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	l = slog.New(handler)
}

func Dubugf(c *gin.Context, format string, args ...interface{}) {
	out(c, "DEBUG", format, args...)
}

func Infof(c *gin.Context, format string, args ...interface{}) {
	out(c, "INFO", format, args...)
}

func Errorf(c *gin.Context, format string, args ...interface{}) {
	out(c, "ERROR", format, args...)
}

func Warnf(c *gin.Context, format string, args ...interface{}) {
	out(c, "WARN", format, args...)
}

type Entry struct {
	Message string
	TraceID string
}

func out(c *gin.Context, severity, format string, args ...interface{}) {

	e := Entry{
		Message: fmt.Sprintf(format, args...),
		TraceID: GetRequestTraceID(c),
	}

	switch severity {
	case "INFO":
		l.Info(e.Message, "trace_id", e.TraceID)
	case "ERROR":
		l.Error(e.Message, "trace_id", e.TraceID)
	case "WARN":
		l.Warn(e.Message, "trace_id", e.TraceID)
	case "DEBUG":
		l.Debug(e.Message, "trace_id", e.TraceID)
	default:
		l.Info(e.Message, "trace_id", e.TraceID)
	}
}
