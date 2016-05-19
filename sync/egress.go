package sync

import (
	"net"
	"os"

	"github.com/bahadley/esp/log"
)

const (
	defaultAddr       = "localhost"
	defaultMasterAddr = "localhost"
	defaultDstPort    = "22219"

	envAddr       = "ESP_ADDR"
	envMasterAddr = "ESP_MASTER_ADDR"
	envDstPort    = "ESP_SYNC_PORT"

	chanbufsz = 10
)

var (
	dstAddr *net.UDPAddr
	srcAddr *net.UDPAddr

	SyncChan chan string
)

func Egress() {
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		log.Error.Fatal(err.Error())
	}
	defer conn.Close()

	for {
		msg := <-SyncChan

		var tx []byte
		copy(tx[:], msg)
		log.Info.Printf("Tx(%s): %s", dstAddr, tx)

		_, err = conn.Write(tx)
		if err != nil {
			log.Warning.Println(err.Error())
		}
	}
}

func init() {
	addr := os.Getenv(envAddr)
	if len(addr) == 0 {
		addr = defaultAddr
	}

	masterAddr := os.Getenv(envMasterAddr)
	if len(masterAddr) == 0 {
		masterAddr = defaultMasterAddr
	}

	dstPort := os.Getenv(envDstPort)
	if len(dstPort) == 0 {
		dstPort = defaultDstPort
	}

	var err error
	dstAddr, err = net.ResolveUDPAddr("udp", masterAddr+":"+dstPort)
	if err != nil {
		log.Error.Fatal(err.Error())
	}

	srcAddr, err = net.ResolveUDPAddr("udp", addr+":0")
	if err != nil {
		log.Error.Fatal(err.Error())
	}

	SyncChan = make(chan string, chanbufsz)
}
