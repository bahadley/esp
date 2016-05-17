package operator

import (
	"os"
	"strconv"
	"sync"

	"github.com/bahadley/esp/log"
)

const (
	defaultBufSz = 4
	defaultAggSz = 2

	envWinSz = "ESP_WINDOW_SIZE"
	envAggSz = "ESP_AGGREGATE_SIZE"
)

var (
	// Invariant:  Descending order by SensorTuple.Timestamp
	window []*SensorTuple

	bufSz uint32
	aggSz uint32

	// Used for window modification critical section.
	mutex sync.Mutex
)

func WindowInsert(msg string) error {
	newTuple := new(SensorTuple)

	err := Unmarshal(msg, newTuple)
	if err != nil {
		log.Warning.Printf("Failed to unmarshal tuple: %s", msg)
	} else {
		mutex.Lock()
		{
			if insert(newTuple) && window[bufSz-1] != nil {
				avg := aggregate()
				aggTuple, err := Marshal(newTuple.Sensor, avg)
				if err != nil {
					log.Warning.Printf("Failed to marshal aggregate tuple for sensor: %s",
						newTuple.Sensor)
				} else {
					EgressChan <- aggTuple
				}
			}
		}
		mutex.Unlock()
	}

	return err
}

func insert(tmp *SensorTuple) bool {
	inserted := false

	if window[0] == nil {
		window[0] = tmp
	} else {
		for idx, st := range window {
			if inserted ||
				(!inserted && st != nil && tmp.Timestamp.After(st.Timestamp)) {
				window[idx] = tmp
				tmp = st
				inserted = true
			}
		}
	}

	return inserted
}

func aggregate() float64 {
	sum := 0.0
	for idx := bufSz - aggSz; idx < bufSz; idx++ {
		sum += window[idx].Data
		window[idx] = nil
	}
	return sum / float64(aggSz)
}

func init() {
	envVal := os.Getenv(envWinSz)
	if len(envVal) == 0 {
		bufSz = defaultBufSz
	} else {
		val, err := strconv.Atoi(envVal)
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

	envVal = os.Getenv(envAggSz)
	if len(envVal) == 0 {
		aggSz = defaultAggSz
	} else {
		val, err := strconv.Atoi(envVal)
		if err != nil {
			log.Error.Fatalf("Invalid environment variable: %s",
				envAggSz)
		}
		aggSz = uint32(val)
	}

	if bufSz < aggSz {
		log.Error.Fatalf("Invalid environment variables: %s < %s",
			envWinSz, envAggSz)
	}

	window = make([]*SensorTuple, bufSz)
}
