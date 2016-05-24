package operator

import (
	"os"
	"strconv"

	"github.com/bahadley/esp/log"
)

const (
	defaultBufSz = 4
	defaultAggSz = 2

	envWinSz = "ESP_WINDOW_SIZE"
	envAggSz = "ESP_AGGREGATE_SIZE"
)

func WindowSize() uint32 {
	var bufSz uint32

	env := os.Getenv(envWinSz)
	if len(env) == 0 {
		bufSz = defaultBufSz
	} else {
		val, err := strconv.Atoi(env)
		if err != nil {
			log.Error.Fatalf("Invalid environment variable: %s",
				envWinSz)
		}
		bufSz = uint32(val)
	}

	if bufSz <= 0 {
		log.Error.Fatalf("Invalid environment variable: %s <= 0",
			envWinSz)
	}

	return bufSz
}

func AggregateSize() uint32 {
	var aggSz uint32

	env := os.Getenv(envAggSz)
	if len(env) == 0 {
		aggSz = defaultAggSz
	} else {
		val, err := strconv.Atoi(env)
		if err != nil {
			log.Error.Fatalf("Invalid environment variable: %s",
				envAggSz)
		}
		aggSz = uint32(val)
	}

	if aggSz <= 0 {
		log.Error.Fatalf("Invalid environment variable: %s <= 0",
			envAggSz)
	}

	return aggSz
}
