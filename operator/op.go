package operator

import (
	"sync"
)

const (
	chanbufsz = 10
)

var (
	IngestChan chan string
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

func QueueMsg(msg string) {
	ingestMutex.Lock()
	{
		IngestChan <- msg
	}
	ingestMutex.Unlock()
}

func init() {
	IngestChan = make(chan string, chanbufsz)
	EgressChan = make(chan []byte, chanbufsz)
}
