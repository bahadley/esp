package sink

import (
	"net"

	"github.com/bahadley/esp/log"
	"github.com/bahadley/esp/system"
)

func Ingress() {
	sinkAddr, err := net.ResolveUDPAddr("udp",
		system.NodeAddr()+":"+system.SinkPort())
	if err != nil {
		log.Error.Fatal(err.Error())
	}

	conn, err := net.ListenUDP("udp", sinkAddr)
	if err != nil {
		log.Error.Fatal(err.Error())
	}

	defer conn.Close()

	log.Info.Printf("Listening for aggregation tuples (%s UDP) ...",
		sinkAddr.String())

	buf := make([]byte, system.TupleBufLen(), system.TupleBufCap())
	for {
		n, caddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Warning.Println(err.Error())
			continue
		}

		log.Trace.Printf("Rx(%s): %s", caddr, buf[0:n])
	}
}
