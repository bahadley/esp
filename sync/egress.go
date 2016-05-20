package sync

import (
	"net"

	"github.com/bahadley/esp/log"
	"github.com/bahadley/esp/system"
)

var (
	SyncChan chan []byte
)

func Egress() {
	masterAddr, err := net.ResolveUDPAddr("udp",
		system.MasterAddr()+":"+system.SyncPort())
	if err != nil {
		log.Error.Fatal(err.Error())
	}

	srcAddr, err := net.ResolveUDPAddr("udp",
		system.NodeAddr()+":0")
	if err != nil {
		log.Error.Fatal(err.Error())
	}

	conn, err := net.DialUDP("udp", srcAddr, masterAddr)
	if err != nil {
		log.Error.Fatal(err.Error())
	}

	defer conn.Close()

	for {
		msg := <-SyncChan

		log.Info.Printf("Tx(%s): %s", masterAddr, msg)

		_, err = conn.Write(msg)
		if err != nil {
			log.Warning.Println(err.Error())
		}
	}
}

func init() {
	SyncChan = make(chan []byte, system.ChannelBufSz())
}
