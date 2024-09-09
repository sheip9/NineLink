package main

import (
	"NineLink/internal/server"
	"golang.org/x/sync/errgroup"
	"log"
)

var (
	g errgroup.Group
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	appServer := server.AppServer()
	g.Go(func() error {
		return appServer.ListenAndServe()
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
