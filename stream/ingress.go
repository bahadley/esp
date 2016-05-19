package stream

import (
	"net"

	"github.com/bahadley/esp/log"
	"github.com/bahadley/esp/operator"
	"github.com/bahadley/esp/sync"
	"github.com/bahadley/esp/system"
)

const (
	msgBufLen = 128
	msgBufCap = 1024
)

func Ingress() {
	ingestAddr, err := net.ResolveUDPAddr("udp",
		system.NodeAddr()+":"+system.IngestPort())
	if err != nil {
		log.Error.Fatal(err.Error())
	}

	conn, err := net.ListenUDP("udp", ingestAddr)
	if err != nil {
		log.Error.Fatal(err.Error())
	}

	defer conn.Close()

	log.Info.Printf("Listening for sensor tuples (%s UDP) ...",
		ingestAddr.String())

	buf := make([]byte, msgBufLen, msgBufCap)
	for {
		n, caddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Warning.Println(err.Error())
			continue
		}

		msg := make([]byte, n)
		copy(msg, buf[0:n])

		log.Info.Printf("Rx(%s): %s", caddr, msg)

		if system.Master() {
			operator.QueueMsg(msg)
		} else {
			sync.SyncChan <- msg
		}
	}
}
