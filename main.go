package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bahadley/esp/log"
	"github.com/bahadley/esp/operator"
	"github.com/bahadley/esp/stream"
	"github.com/bahadley/esp/sync"
)

func main() {
	log.Info.Println("Starting up ...")

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		// Block waiting for signal.
		<-c
		log.Info.Println("Shutting down ...")
		os.Exit(0)
	}()

	go stream.Egress()
	go sync.Ingress()
	go operator.Ingest()
	stream.Ingress()
}
