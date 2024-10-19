package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sheip9/ninelink/internal/constant"
)

func WithHeaderDealing() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Powered-By", "NineLink"+constant.Version)
		c.Next()
	}
}
