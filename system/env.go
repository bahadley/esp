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

	masterFlag        = "YES"
	defaultNodeAddr   = "localhost"
	defaultSinkAddr   = "localhost"
	defaultMasterAddr = "localhost"
	defaultIngestPort = "22221"
	defaultSinkPort   = "22220"
	defaultSyncPort   = "22219"
)

var (
	Master bool

	NodeAddr string
	SinkAddr string
	// MasterAddr is used by nonmaster nodes to reach the master node.
	MasterAddr string

	IngestPort string
	SinkPort   string
	SyncPort   string
)

func init() {
	m := os.Getenv(envMaster)
	if len(m) > 0 && strings.ToUpper(m) == masterFlag {
		Master = true
	}

	NodeAddr := os.Getenv(envNodeAddr)
	if len(NodeAddr) == 0 {
		NodeAddr = defaultNodeAddr
	}

	SinkAddr := os.Getenv(envSinkAddr)
	if len(SinkAddr) == 0 {
		SinkAddr = defaultSinkAddr
	}

	MasterAddr := os.Getenv(envMasterAddr)
	if len(MasterAddr) == 0 {
		MasterAddr = defaultMasterAddr
	}

	IngestPort := os.Getenv(envIngestPort)
	if len(IngestPort) == 0 {
		IngestPort = defaultIngestPort
	}

	SinkPort := os.Getenv(envSinkPort)
	if len(SinkPort) == 0 {
		SinkPort = defaultSinkPort
	}

	SyncPort := os.Getenv(envSyncPort)
	if len(SyncPort) == 0 {
		SyncPort = defaultSyncPort
	}
}
