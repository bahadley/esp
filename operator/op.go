package operator

const (
	chanbufsz = 10
)

var (
	IngestChan chan string
	EgressChan chan []byte
)

func Ingest() {
	for {
		msg := <-IngestChan
		_ = WindowAppend(msg)
	}
}

func init() {
	IngestChan = make(chan string, chanbufsz)
	EgressChan = make(chan []byte, chanbufsz)
}
