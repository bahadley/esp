package operator

import (
	"os"
	"strconv"
	"sync"

	"github.com/bahadley/esp/log"
)

const (
	defaultLength  = 4
	defaultTrigger = 2

	envWinLen  = "ESP_WINDOW_LENGTH"
	envWinTrig = "ESP_WINDOW_TRIGGER"
)

var (
	// Invariant:  Descending order by SensorTuple.Timestamp
	window []*SensorTuple

	length  uint32
	trigger uint32

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
			if insert(newTuple) && window[length-1] != nil {
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
	for idx := length - trigger; idx < length; idx++ {
		sum += window[idx].Data
		window[idx] = nil
	}
	return sum / float64(trigger)
}

func init() {
	envVal := os.Getenv(envWinLen)
	if len(envVal) == 0 {
		length = defaultLength
	} else {
		val, err := strconv.Atoi(envVal)
		if err != nil {
			log.Error.Fatalf("Invalid environment variable: %s",
				envWinLen)
		}
		length = uint32(val)
	}

	if length <= 0 {
		log.Error.Fatalf("Invalid environment variable: %s <= 0",
			envWinLen)
	}

	envVal = os.Getenv(envWinTrig)
	if len(envVal) == 0 {
		trigger = defaultTrigger
	} else {
		val, err := strconv.Atoi(envVal)
		if err != nil {
			log.Error.Fatalf("Invalid environment variable: %s",
				envWinTrig)
		}
		trigger = uint32(val)
	}

	if length < trigger {
		log.Error.Fatalf("Invalid environment variables: %s < %s",
			envWinLen, envWinTrig)
	}

	window = make([]*SensorTuple, length)
}
