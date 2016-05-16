package telemetry

import (
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
	udpAddr *net.UDPAddr
)

func Ingest() {
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Error.Fatal(err.Error())
	}
	defer conn.Close()

	log.Info.Printf("Listening for sensor tuples (%s UDP) ...",
		udpAddr.String())

	op := make(chan string, chanbufsz)
	go operator.Ingest(op)

	buf := make([]byte, 1024)
	for {
		n, caddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Warning.Println(err.Error())
			continue
		}

		msg := string(buf[0:n])
		log.Info.Printf("Recv(%s): %s", caddr, msg)
		op <- msg
	}
}

func init() {
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
		log.Error.Fatal(err.Error())
	}
}
