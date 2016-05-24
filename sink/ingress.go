package sink

import (
	"fmt"
	"net"
	"time"

	"github.com/bahadley/esp/log"
	"github.com/bahadley/esp/operator"
	"github.com/bahadley/esp/system"
)

var (
	outputChan chan *operator.SensorTuple
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

		arrivalTime := time.Now().UnixNano()

		log.Trace.Printf("Rx(%s): %s", caddr, buf[0:n])

		aggTuple, err := operator.Unmarshal(buf[0:n])
		if err != nil {
			log.Warning.Printf("Failed to unmarshal tuple: %s", buf[0:n])
			continue
		}

		aggTuple.Timestamp = arrivalTime

		outputChan <- aggTuple
	}
}

func Output() {
	outputChan = make(chan *operator.SensorTuple, system.ChannelBufSz())

	for {
		aggTuple := <-outputChan

		fmt.Printf("%d,%s,%.2f\n", aggTuple.Timestamp,
			aggTuple.Type, aggTuple.Data)
	}
}
