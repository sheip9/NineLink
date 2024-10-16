package server

import (
	"fmt"
	"github.com/sheip9/ninelink/config"
	"github.com/sheip9/ninelink/internal/router"
	"net/http"
	"time"
)

var (
	c = config.Conf
)

func AppServer() *http.Server {
	h := router.AppRouter()
	h.LoadHTMLGlob("./web/template/*")
	return &http.Server{
		Addr:         fmt.Sprintf("%s:%d", (*c).IP, (*c).Port),
		Handler:      h,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
