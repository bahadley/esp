package operator

import (
	"encoding/json"
	"net"
	"os"
	"time"

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
	dstAddr *net.UDPAddr
	srcAddr *net.UDPAddr
)

type (
	gMsg struct {
		Sensor    string    `json:"sensor"`
		Type      string    `json:"type"`
		Timestamp time.Time `json:"ts"`
		Data      float64   `json:"data"`
	}
)

func Window(ingest chan string) {
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		log.Logoutput(log.ErrPrefix, err.Error())
	}
	defer conn.Close()

	var gm gMsg
	sum := 0.0
	for count := 0; ; count++ {
		msg := <-ingest
		err := json.Unmarshal([]byte(msg), &gm)
		if err != nil {
			log.Logoutput(log.ErrPrefix, err.Error())
			continue
		}

		sum += gm.Data

		if count%2 == 1 {
			gm.Type = "TA"
			gm.Data = sum / 2.0
			data, err := json.Marshal(gm)
			if err != nil {
				log.Logoutput(log.ErrPrefix, err.Error())
			}

			buf := []byte(data)
			_, err = conn.Write(buf)
			if err != nil {
				log.Logoutput(log.ErrPrefix, err.Error())
			}

			count = -1
			sum = 0.0
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
		log.Logfatalerror(err)
	}

	srcAddr, err = net.ResolveUDPAddr("udp", addr+":0")
	if err != nil {
		log.Logfatalerror(err)
	}
}
