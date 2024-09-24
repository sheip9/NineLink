package main

import (
	"flag"
	"github.com/sheip9/ninelink/config"
	"github.com/sheip9/ninelink/internal/server"
	"github.com/sheip9/ninelink/internal/utils"
	"golang.org/x/sync/errgroup"
	"log"
)

var (
	g errgroup.Group
)

func main() {
	flag.StringVar(&config.File, "c", "./config.yml", "Path to config file")
	flag.Parse()
	startInitApp()

	appServer := server.AppServer()
	g.Go(func() error {
		return appServer.ListenAndServe()
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
func startInitApp() {
	config.ReadConfig()
	utils.InitDB()
}
