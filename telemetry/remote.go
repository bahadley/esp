package telemetry

import (
	"fmt"
	"net"
	"os"

	"github.com/bahadley/esp/log"
	"github.com/bahadley/esp/operator"
)

const (
	defaultAddr = "localhost"
	defaultPort = "22221"

	envsaddr = "ESP_ADDR"
	envport  = "ESP_PORT"

	chanbufsz = 10
)

var (
	op chan string

	udpAddr *net.UDPAddr
)

func Ingest() {
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Logoutput(log.ErrPrefix, err.Error())
	}
	defer conn.Close()

	log.Logoutput(log.InfoPrefix,
		fmt.Sprintf("Listening for sensor tuples (%s UDP) ...", udpAddr.String()))

	buf := make([]byte, 1024)
	for {
		n, caddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Logoutput(log.ErrPrefix, err.Error())
			continue
		}

		msg := string(buf[0:n])
		log.Logoutput(log.InfoPrefix,
			fmt.Sprintf("Recv(%s): %s", caddr, msg))
		op <- msg
	}
}

func init() {
	// Launch operator goroutine and establish channel to it.
	op = make(chan string, chanbufsz)
	go operator.Window(op)

	// Build the UDP address that we will listen on.
	addr := os.Getenv(envsaddr)
	if len(addr) == 0 {
		addr = defaultAddr
	}

	port := os.Getenv(envport)
	if len(port) == 0 {
		port = defaultPort
	}

	var err error
	udpAddr, err = net.ResolveUDPAddr("udp", addr+":"+port)
	if err != nil {
		log.Logfatalerror(err)
	}
}
