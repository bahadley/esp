package stream

import (
	"net"

	"github.com/bahadley/esp/log"
	"github.com/bahadley/esp/operator"
	"github.com/bahadley/esp/system"
)

func Egress() {
	sinkAddr, err := net.ResolveUDPAddr("udp",
		system.SinkAddr()+":"+system.SinkPort())
	if err != nil {
		log.Error.Fatal(err.Error())
	}

	srcAddr, err := net.ResolveUDPAddr("udp",
		system.NodeAddr()+":0")
	if err != nil {
		log.Error.Fatal(err.Error())
	}

	conn, err := net.DialUDP("udp", srcAddr, sinkAddr)
	if err != nil {
		log.Error.Fatal(err.Error())
	}

	defer conn.Close()

	for {
		msg := <-operator.EgressChan

		log.Info.Printf("Tx(%s): %s", sinkAddr, msg)

		_, err = conn.Write(msg)
		if err != nil {
			log.Warning.Println(err.Error())
		}
	}
}
