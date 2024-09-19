package middleware

import (
	"inbody-ocr-backend/internal/domain/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(st service.TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		valid, err := st.TokenValid(c)
		if err != nil || !valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Next()
	}
}
