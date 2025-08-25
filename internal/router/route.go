package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sheip9/ninelink/internal/handler/app"
	"github.com/sheip9/ninelink/internal/middleware"
)

// AppRouter 创建并配置应用路由
func AppRouter() *gin.Engine {
	r := gin.Default()

	// 使用中间件
	r.Use(middleware.WithHeaderDealing)

	// 注册路由
	r.GET("/", app.GetIndex)
	r.GET("/:path", app.GetRecord)

	return r
}
