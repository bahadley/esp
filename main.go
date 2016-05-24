package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bahadley/esp/log"
	"github.com/bahadley/esp/operator"
	"github.com/bahadley/esp/sink"
	"github.com/bahadley/esp/stream"
	"github.com/bahadley/esp/sync"
	"github.com/bahadley/esp/system"
)

func main() {
	log.Info.Println("Starting up ...")

	// Allow the node to be shut down gracefully.
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		// Block waiting for signal.
		<-c
		log.Info.Println("Shutting down ...")
		os.Exit(0)
	}()

	if system.Master() {
		go stream.Egress()
		go sync.Ingress()
		go operator.Ingest()
		stream.Ingress()
	} else if system.Sink() {
		sink.Ingress()
	} else {
		// Non-master node
		go sync.Egress()
		stream.Ingress()
	}
}
