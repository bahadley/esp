package main

import (
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bahadley/esp/log"
	"github.com/bahadley/esp/operator"
	"github.com/bahadley/esp/stream"
	"github.com/bahadley/esp/sync"
)

const (
	envMaster = "ESP_MASTER"

	masterFlag = "NO"
)

var (
	master bool = true
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

	if master {
		go stream.Egress()
		go sync.Ingress()
		go operator.Ingest()
		stream.Ingress(true)
	} else {
		go sync.Egress()
		stream.Ingress(false)
	}
}

func init() {
	m := os.Getenv(envMaster)
	if len(m) == 0 && strings.ToUpper(m) == masterFlag {
		master = false
	}
}
