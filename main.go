package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bahadley/esp/log"
	"github.com/bahadley/esp/telemetry"
)

func main() {
	log.Logoutput(log.InfoPrefix, "Starting up ...")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		// Block waiting for signal.
		<-c
		log.Logoutput(log.InfoPrefix, "Shutting down ...")
		os.Exit(0)
	}()

	telemetry.Ingest()
}
