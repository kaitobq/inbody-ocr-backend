package render

import (
	"encoding/json"
	"inbody-ocr-backend/internal/infra/logging"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  int    `json:"status"`
}

func ErrorJSON(c *gin.Context, err string, status int) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(status)

	if err := json.NewEncoder(c.Writer).Encode(&Error{
		Message: err,
		Status:  status,
	}); err != nil {
		logging.Warnf(c, "ErrorJSON json.NewEncoder %v", err)
	}
}

func ErrorCodeJSON(c *gin.Context, err string, status, code int) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(status)

	if err := json.NewEncoder(c.Writer).Encode(&Error{
		Message: err,
		Code:    code,
		Status:  status,
	}); err != nil {
		logging.Warnf(c, "ErrorCodeJSON json.NewEncoder %v", err)
	}
}
