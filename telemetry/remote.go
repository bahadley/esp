package telemetry

import (
	"fmt"
	"net"
	"os"

	"github.com/bahadley/esp/log"
)

const (
	defaultAddr = "127.0.0.1"
	defaultPort = "22221"

	envsaddr = "ESP_ADDR"
	envport  = "ESP_PORT"
)

func Ingest() {
	addr := os.Getenv(envsaddr)
	if len(addr) == 0 {
		addr = defaultAddr
	}

	port := os.Getenv(envport)
	if len(port) == 0 {
		port = defaultPort
	}

	saddr, err := net.ResolveUDPAddr("udp", addr+":"+port)
	if err != nil {
		log.Logfatalerror(err)
	}

	conn, err := net.ListenUDP("udp", saddr)
	if err != nil {
		log.Logoutput(log.ErrPrefix, err.Error())
	}
	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		n, caddr, err := conn.ReadFromUDP(buf)
		log.Logoutput(log.InfoPrefix,
			fmt.Sprintf("Recv(%s): %s", caddr, string(buf[0:n])))

		if err != nil {
			log.Logoutput(log.ErrPrefix, err.Error())
		}
	}
}
