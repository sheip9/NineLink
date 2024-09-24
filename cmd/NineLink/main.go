package main

import (
	"github.com/sheip9/ninelink/internal/server"
	"golang.org/x/sync/errgroup"
	"log"
)

var (
	g errgroup.Group
)

func main() {
	appServer := server.AppServer()
	g.Go(func() error {
		return appServer.ListenAndServe()
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
