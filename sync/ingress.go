package sync

import (
	"net"

	"github.com/bahadley/esp/log"
	"github.com/bahadley/esp/operator"
	"github.com/bahadley/esp/system"
)

func Ingress() {
	syncAddr, err := net.ResolveUDPAddr("udp",
		system.NodeAddr()+":"+system.SyncPort())
	if err != nil {
		log.Error.Fatal(err.Error())
	}

	conn, err := net.ListenUDP("udp", syncAddr)
	if err != nil {
		log.Error.Fatal(err.Error())
	}

	defer conn.Close()

	log.Info.Printf("Listening for synchronization tuples (%s UDP) ...",
		syncAddr.String())

	buf := make([]byte, system.TupleBufLen(), system.TupleBufCap())
	for {
		n, caddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Warning.Println(err.Error())
			continue
		}

		msg := make([]byte, n)
		copy(msg, buf[0:n])

		log.Trace.Printf("Sync(%s): %s", caddr, msg)

		operator.QueueMsg(msg)
	}
}
