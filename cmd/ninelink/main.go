package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sheip9/ninelink/config"
	"github.com/sheip9/ninelink/internal/database"
	"github.com/sheip9/ninelink/internal/server"
	"golang.org/x/sync/errgroup"
	"log"
)

func main() {
	flag.StringVar(&config.File, "c", "./config.yml", "Path to config file")
	flag.Parse()

	if err := initApp(); err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	gin.SetMode(config.Conf.GetGinMode())
	appServer := server.AppServer()

	var g errgroup.Group
	g.Go(func() error {
		if err := appServer.ListenAndServe(); err != nil {
			return fmt.Errorf("server error: %v", err)
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

// initApp 初始化应用程序
func initApp() error {
	config.InitConfig()

	database.InitDB()

	return nil
}
