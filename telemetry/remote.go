package telemetry

import (
	"fmt"
	"net"
	"os"

	"github.com/bahadley/esp/log"
	"github.com/bahadley/esp/operator"
)

const (
	defaultAddr = "127.0.0.1"
	defaultPort = "22221"

	envsaddr = "ESP_ADDR"
	envport  = "ESP_PORT"
)

var (
    op chan string
)

func Ingest() {
	saddr, err := resolveAddr()
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
		if err != nil {
			log.Logoutput(log.ErrPrefix, err.Error())
            continue
		}

        msg := string(buf[0:n])
		log.Logoutput(log.InfoPrefix, fmt.Sprintf("Recv(%s): %s", caddr, msg))
        op <- msg
	}
}

func init() {
    op = make(chan string, 10)
    go operator.Window(op)
}

func resolveAddr() (*net.UDPAddr, error) {
	addr := os.Getenv(envsaddr)
	if len(addr) == 0 {
		addr = defaultAddr
	}

	port := os.Getenv(envport)
	if len(port) == 0 {
		port = defaultPort
	}

	return net.ResolveUDPAddr("udp", addr + ":" + port)
}
