package router

import (
	"NineLink/internal/handler/app"
	"NineLink/internal/middleware"
	"github.com/gin-gonic/gin"
)

func AppRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.WithHeaderDealing())
	r.GET("/", app.GetIndex)
	r.GET("/:path", app.GetRecord)
	return r
}
