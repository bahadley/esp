package system

import (
	"os"
	"strings"
)

const (
	envNodeType   = "ESP_NODE_TYPE"
	envNodeAddr   = "ESP_ADDR"
	envSinkAddr   = "ESP_SINK_ADDR"
	envMasterAddr = "ESP_MASTER_ADDR"
	envIngestPort = "ESP_PORT"
	envSinkPort   = "ESP_SINK_PORT"
	envSyncPort   = "ESP_SYNC_PORT"
	envTrace      = "ESP_TRACE"

	nonMasterNode = "N"
	masterNode    = "M"
	sinkNode      = "S"

	defaultNodeType    = nonMasterNode
	defaultNodeAddr    = "localhost"
	defaultSinkAddr    = "localhost"
	defaultMasterAddr  = "localhost"
	defaultIngestPort  = "22221"
	defaultSinkPort    = "22220"
	defaultSyncPort    = "22219"
	traceFlag          = "YES"
	defaultTupleBufLen = 128
	defaultTupleBufCap = 1024
	defaultChanBufSz   = 10
)

var (
	nodeType string
)

func Master() bool {
	if len(nodeType) == 0 {
		m := os.Getenv(envNodeType)
		if len(m) > 0 && strings.ToUpper(m) == masterNode {
			nodeType = masterNode
		}
	}

	return nodeType == masterNode
}

func Sink() bool {
	if len(nodeType) == 0 {
		s := os.Getenv(envNodeType)
		if len(s) > 0 && strings.ToUpper(s) == sinkNode {
			nodeType = sinkNode
		}
	}

	return nodeType == sinkNode
}

func NodeAddr() string {
	addr := os.Getenv(envNodeAddr)
	if len(addr) == 0 {
		return defaultNodeAddr
	} else {
		return addr
	}
}

func SinkAddr() string {
	addr := os.Getenv(envSinkAddr)
	if len(addr) == 0 {
		return defaultSinkAddr
	} else {
		return addr
	}
}

func MasterAddr() string {
	addr := os.Getenv(envMasterAddr)
	if len(addr) == 0 {
		return defaultMasterAddr
	} else {
		return addr
	}
}

func IngestPort() string {
	port := os.Getenv(envIngestPort)
	if len(port) == 0 {
		return defaultIngestPort
	} else {
		return port
	}
}

func SinkPort() string {
	port := os.Getenv(envSinkPort)
	if len(port) == 0 {
		return defaultSinkPort
	} else {
		return port
	}
}

func SyncPort() string {
	port := os.Getenv(envSyncPort)
	if len(port) == 0 {
		return defaultSyncPort
	} else {
		return port
	}
}

func Trace() bool {
	t := os.Getenv(envTrace)
	if len(t) > 0 && strings.ToUpper(t) == traceFlag {
		return true
	} else {
		return false
	}
}

func TupleBufLen() uint32 {
	return defaultTupleBufLen
}

func TupleBufCap() uint32 {
	return defaultTupleBufCap
}

func ChannelBufSz() int {
	return defaultChanBufSz
}
