package system

import (
	"os"
	"strings"
)

const (
	envMaster     = "ESP_MASTER"
	envNodeAddr   = "ESP_ADDR"
	envSinkAddr   = "ESP_SINK_ADDR"
	envMasterAddr = "ESP_MASTER_ADDR"
	envIngestPort = "ESP_PORT"
	envSinkPort   = "ESP_SINK_PORT"
	envSyncPort   = "ESP_SYNC_PORT"

	masterFlag         = "YES"
	defaultNodeAddr    = "localhost"
	defaultSinkAddr    = "localhost"
	defaultMasterAddr  = "localhost"
	defaultIngestPort  = "22221"
	defaultSinkPort    = "22220"
	defaultSyncPort    = "22219"
	defaultTupleBufLen = 128
	defaultTupleBufCap = 1024
	defaultChanBufSz   = 10
)

func Master() bool {
	m := os.Getenv(envMaster)
	if len(m) > 0 && strings.ToUpper(m) == masterFlag {
		return true
	} else {
		return false
	}
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

func TupleBufLen() uint32 {
	return defaultTupleBufLen
}

func TupleBufCap() uint32 {
	return defaultTupleBufCap
}

func ChannelBufSz() int {
	return defaultChanBufSz
}
