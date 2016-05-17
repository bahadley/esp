package operator

import (
	"net"
	"os"

	"github.com/bahadley/esp/log"
)

const (
	defaultAddr     = "localhost"
	defaultSinkAddr = "localhost"
	defaultSinkPort = "22220"

	envsaddr    = "ESP_ADDR"
	envsinkaddr = "ESP_SINK_ADDR"
	envsinkport = "ESP_SINK_PORT"

	chanbufsz = 10
)

var (
	IngestChan chan string

	dstAddr *net.UDPAddr
	srcAddr *net.UDPAddr
)

func Ingest() {
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		log.Error.Fatal(err.Error())
	}
	defer conn.Close()

	for {
		msg := <-IngestChan

		rslt := WindowInsert(msg)
		if len(rslt) > 0 {
			_, err = conn.Write(rslt)
			if err != nil {
				log.Warning.Println(err.Error())
			}
		}
	}
}

func init() {
	addr := os.Getenv(envsaddr)
	if len(addr) == 0 {
		addr = defaultAddr
	}

	sinkAddr := os.Getenv(envsinkaddr)
	if len(sinkAddr) == 0 {
		sinkAddr = defaultSinkAddr
	}

	sinkPort := os.Getenv(envsinkport)
	if len(sinkPort) == 0 {
		sinkPort = defaultSinkPort
	}

	var err error
	dstAddr, err = net.ResolveUDPAddr("udp", sinkAddr+":"+sinkPort)
	if err != nil {
		log.Error.Fatal(err.Error())
	}

	srcAddr, err = net.ResolveUDPAddr("udp", addr+":0")
	if err != nil {
		log.Error.Fatal(err.Error())
	}

	IngestChan = make(chan string, chanbufsz)
}
