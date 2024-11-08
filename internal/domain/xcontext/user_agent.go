package xcontext

import "github.com/gin-gonic/gin"

var (
	uaStrKey = "UA:STR"
)

func WithUserAgent(c *gin.Context, uaStr string) {
	c.Set(uaStrKey, uaStr)
}

func UserAgent(c *gin.Context) string {
	if v, ok := c.Get(uaStrKey); ok {
		return v.(string)
	}
	return ""
}
