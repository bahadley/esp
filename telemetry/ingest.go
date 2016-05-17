package telemetry

import (
	"net"
	"os"

	"github.com/bahadley/esp/log"
	"github.com/bahadley/esp/operator"
)

const (
	defaultIngestAddr = "localhost"
	defaultIngestPort = "22221"

	envIngestAddr = "ESP_ADDR"
	envIngestPort = "ESP_PORT"
)

var (
	IngestAddr *net.UDPAddr
)

func Ingest() {
	conn, err := net.ListenUDP("udp", IngestAddr)
	if err != nil {
		log.Error.Fatal(err.Error())
	}
	defer conn.Close()

	log.Info.Printf("Listening for sensor tuples (%s UDP) ...",
		IngestAddr.String())

	go operator.Ingest()

	buf := make([]byte, 1024)
	for {
		n, caddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Warning.Println(err.Error())
			continue
		}

		msg := string(buf[0:n])
		log.Info.Printf("Recv(%s): %s", caddr, msg)
		operator.IngestChan <- msg
	}
}

func init() {
	// Build the UDP address that we will listen on.
	addr := os.Getenv(envIngestAddr)
	if len(addr) == 0 {
		addr = defaultIngestAddr
	}

	port := os.Getenv(envIngestPort)
	if len(port) == 0 {
		port = defaultIngestPort
	}

	var err error
	IngestAddr, err = net.ResolveUDPAddr("udp", addr+":"+port)
	if err != nil {
		log.Error.Fatal(err.Error())
	}
}
