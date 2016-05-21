package operator

import (
	"github.com/bahadley/esp/system"
)

var (
	IngestChan chan []byte
	EgressChan chan []byte
)

func QueueMsg(msg []byte) {
	IngestChan <- msg
}

func Ingest() {
	for {
		msg := <-IngestChan
		windowInsert(msg)
	}
}

func init() {
	bufSz := system.ChannelBufSz()
	IngestChan = make(chan []byte, bufSz)
	EgressChan = make(chan []byte, bufSz)
}
