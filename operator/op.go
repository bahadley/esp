package operator

import (
	"sync"
)

const (
	chanbufsz = 10
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
		_ = WindowInsert(msg)
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
	IngestChan = make(chan []byte, chanbufsz)
	EgressChan = make(chan []byte, chanbufsz)
}
