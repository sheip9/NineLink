package server

import (
	"fmt"
	"github.com/sheip9/ninelink/config"
	"github.com/sheip9/ninelink/internal/router"
	"net/http"
	"time"
)

func AppServer() *http.Server {
	h := router.AppRouter()
	h.LoadHTMLGlob("./web/template/*")

	// 获取配置
	conf := config.Conf

	return &http.Server{
		Addr:         fmt.Sprintf("%s:%d", conf.IP, conf.Port),
		Handler:      h,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
}
