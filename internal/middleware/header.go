package middleware

import (
	"github.com/gin-gonic/gin"
)

func WithHeaderDealing() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Powered-By", "NineLink")
		c.Next()
	}
}
