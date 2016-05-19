package stream

import (
	"net"
	"os"

	"github.com/bahadley/esp/log"
	"github.com/bahadley/esp/operator"
	"github.com/bahadley/esp/sync"
)

const (
	defaultIngressAddr = "localhost"
	defaultIngressPort = "22221"

	envIngressAddr = "ESP_ADDR"
	envIngressPort = "ESP_PORT"

	msgBufLen = 128
	msgBufCap = 1024
)

var (
	IngressAddr *net.UDPAddr
)

func Ingress(master bool) {
	conn, err := net.ListenUDP("udp", IngressAddr)
	if err != nil {
		log.Error.Fatal(err.Error())
	}
	defer conn.Close()

	log.Info.Printf("Listening for sensor tuples (%s UDP) ...",
		IngressAddr.String())

	buf := make([]byte, msgBufLen, msgBufCap)
	for {
		n, caddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Warning.Println(err.Error())
			continue
		}

		msg := string(buf[0:n])
		log.Info.Printf("Rx(%s): %s", caddr, msg)

		if master {
			operator.QueueMsg(msg)
		} else {
			sync.SyncChan <- msg
		}
	}
}

func init() {
	// Build the UDP address that we will listen on.
	addr := os.Getenv(envIngressAddr)
	if len(addr) == 0 {
		addr = defaultIngressAddr
	}

	port := os.Getenv(envIngressPort)
	if len(port) == 0 {
		port = defaultIngressPort
	}

	var err error
	IngressAddr, err = net.ResolveUDPAddr("udp", addr+":"+port)
	if err != nil {
		log.Error.Fatal(err.Error())
	}
}
