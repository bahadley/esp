package stream

import (
	"net"
	"os"

	"github.com/bahadley/esp/log"
	"github.com/bahadley/esp/operator"
)

const (
	defaultSyncAddr = "localhost"
	defaultSyncPort = "22219"

	envSyncAddr = "ESP_ADDR"
	envSyncPort = "ESP_SYNC_PORT"

	msgBufLen = 128
	msgBufCap = 1024
)

var (
	SyncAddr *net.UDPAddr
)

func Ingress() {
	conn, err := net.ListenUDP("udp", SyncAddr)
	if err != nil {
		log.Error.Fatal(err.Error())
	}
	defer conn.Close()

	log.Info.Printf("Listening for synchronization tuples (%s UDP) ...",
		IngressAddr.String())

	buf := make([]byte, msgBufLen, msgBufCap)
	for {
		n, caddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Warning.Println(err.Error())
			continue
		}

		msg := string(buf[0:n])
		log.Info.Printf("Sync(%s): %s", caddr, msg)
		operator.QueueMsg(msg)
	}
}

func init() {
	// Build the UDP address that we will listen on.
	addr := os.Getenv(envSyncAddr)
	if len(addr) == 0 {
		addr = defaultSyncAddr
	}

	port := os.Getenv(envSyncPort)
	if len(port) == 0 {
		port = defaultSyncPort
	}

	var err error
	SyncAddr, err = net.ResolveUDPAddr("udp", addr+":"+port)
	if err != nil {
		log.Error.Fatal(err.Error())
	}
}
