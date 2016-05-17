package stream

import (
	"net"
	"os"

	"github.com/bahadley/esp/log"
	"github.com/bahadley/esp/operator"
)

const (
	defaultEgressAddr = "localhost"
	defaultSinkAddr   = "localhost"
	defaultSinkPort   = "22220"

	envEgressAddr = "ESP_ADDR"
	envSinkAddr   = "ESP_SINK_ADDR"
	envSinkPort   = "ESP_SINK_PORT"
)

var (
	dstAddr *net.UDPAddr
	srcAddr *net.UDPAddr
)

func Egress() {
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		log.Error.Fatal(err.Error())
	}
	defer conn.Close()

	for {
		msg := <-operator.EgressChan

		log.Info.Printf("Tx(%s): %s", dstAddr, msg)

		_, err = conn.Write(msg)
		if err != nil {
			log.Warning.Println(err.Error())
		}
	}
}

func init() {
	addr := os.Getenv(envEgressAddr)
	if len(addr) == 0 {
		addr = defaultEgressAddr
	}

	sinkAddr := os.Getenv(envSinkAddr)
	if len(sinkAddr) == 0 {
		sinkAddr = defaultSinkAddr
	}

	sinkPort := os.Getenv(envSinkPort)
	if len(sinkPort) == 0 {
		sinkPort = defaultSinkPort
	}

	var err error
	dstAddr, err = net.ResolveUDPAddr("udp", sinkAddr+":"+sinkPort)
	if err != nil {
		log.Error.Fatal(err.Error())
	}

	srcAddr, err = net.ResolveUDPAddr("udp", addr+":0")
	if err != nil {
		log.Error.Fatal(err.Error())
	}
}
