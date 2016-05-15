package operator

import (
	"encoding/json"
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
)

var (
	dstAddr *net.UDPAddr
	srcAddr *net.UDPAddr

	window []SensorTuple
	count  int
)

func Ingest(ingest chan string) {
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		log.Error.Fatal(err.Error())
	}
	defer conn.Close()

	var st SensorTuple
	for {
		msg := <-ingest
		err := json.Unmarshal([]byte(msg), &st)
		if err != nil {
			log.Warning.Println(err.Error())
			continue
		}

		val, threshold := appendTuple(st)
		if threshold {
			st.Type = "TA"
			st.Data = val
			data, err := json.Marshal(st)
			if err != nil {
				log.Warning.Println(err.Error())
			}

			buf := []byte(data)
			_, err = conn.Write(buf)
			if err != nil {
				log.Warning.Println(err.Error())
			}
		}
	}
}

func appendTuple(st SensorTuple) (float64, bool) {
	var tmp SensorTuple
	for idx, val := range window {
		if idx == 0 {
			window[idx] = st
		} else {
			window[idx] = tmp
		}
		tmp = val
	}

	count++
	if count == 2 {
		count = 0
		return (window[0].Data + window[1].Data) / 2.0, true
	}
	return 0.0, false
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

	window = make([]SensorTuple, 4)
}
