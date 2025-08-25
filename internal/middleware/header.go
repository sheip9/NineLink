package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sheip9/ninelink/internal/constant"
)

func WithHeaderDealing(c *gin.Context) {
	c.Header("X-Powered-By", "NineLink"+constant.GetVersion())
	c.Next()
}
