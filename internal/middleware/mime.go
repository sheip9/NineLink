package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sheip9/ninelink/internal/enum"
	"strings"
)

func AcceptHeader(c *gin.Context) {
	if strings.Contains(c.GetHeader("Accept"), "application/json") {
		c.Set(enum.RespType, enum.JSON)
	} else if strings.Contains(c.GetHeader("Accept"), "text/html") {
		c.Set(enum.RespType, enum.HTML)
	}
	c.Next()
}
