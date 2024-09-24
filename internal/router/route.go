package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sheip9/ninelink/internal/handler/app"
	"github.com/sheip9/ninelink/internal/middleware"
)

func AppRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.WithHeaderDealing())
	r.GET("/", app.GetIndex)
	r.GET("/:path", app.GetRecord)
	return r
}
