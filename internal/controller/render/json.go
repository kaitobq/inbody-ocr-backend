package render

import (
	"encoding/json"
	"inbody-ocr-backend/internal/infra/logging"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSON(c *gin.Context, res interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(c.Writer).Encode(res); err != nil {
		logging.Warnf(c, "JSON json.NewEncoder %v", err)
	}
}

func OK(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
}
