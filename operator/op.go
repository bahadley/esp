package operator

import (
	"sync"

	"github.com/bahadley/esp/system"
)

var (
	IngestChan chan []byte
	EgressChan chan []byte

	// Used for IngestChan write critical section.
	ingestMutex sync.Mutex
)

func Ingest() {
	for {
		msg := <-IngestChan
		windowInsert(msg)
	}
}

func QueueMsg(msg []byte) {
	ingestMutex.Lock()
	{
		IngestChan <- msg
	}
	ingestMutex.Unlock()
}

func init() {
	bufSz := system.ChannelBufSz()
	IngestChan = make(chan []byte, bufSz)
	EgressChan = make(chan []byte, bufSz)
}
