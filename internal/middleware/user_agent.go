package middleware

import (
	"fmt"
	"inbody-ocr-backend/internal/domain/xcontext"

	"github.com/gin-gonic/gin"
)

type UserAgent struct{}

func NewUserAgent() *UserAgent {
	return &UserAgent{}
}

func (u *UserAgent) UserAgent() gin.HandlerFunc {
	return func(c *gin.Context) {
		ua := c.Request.UserAgent()
		xcontext.WithUserAgent(c, ua)
		fmt.Println("User-Agent:", ua)

		c.Next()
	}
}
