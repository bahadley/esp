package operator

import (
	"os"
	"strconv"

	"github.com/bahadley/esp/log"
)

const (
	defaultCapacity   = 4
	defaultTriggerQty = 2

	envWinCap  = "ESP_WINDOW_CAPACITY"
	envWinTrig = "ESP_WINDOW_TRIGGER"
)

var (
	window    []SensorTuple
	count     int
	trigger   int
	avgFactor float64
)

func AppendTuple(st SensorTuple) (float64, bool) {
	var tmp SensorTuple
	for idx, val := range window {
		if idx == 0 {
			window[idx] = st
		} else {
			window[idx] = tmp
		}
		tmp = val
	}

	count++
	if count == trigger {
		count = 0
		return (window[0].Data + window[1].Data) / avgFactor, true
	}
	return 0.0, false
}

func init() {
	var err error
	var cap int

	envVal := os.Getenv(envWinCap)
	if len(envVal) == 0 {
		cap = defaultCapacity
	} else {
		cap, err = strconv.Atoi(envVal)
		if err != nil {
			log.Error.Fatalf("Invalid environment variable: %s", envWinCap)
		}
	}

	envVal = os.Getenv(envWinTrig)
	if len(envVal) == 0 {
		trigger = defaultTriggerQty
	} else {
		trigger, err = strconv.Atoi(envVal)
		if err != nil {
			log.Error.Fatalf("Invalid environment variable: %s", envWinTrig)
		}
	}

	window = make([]SensorTuple, cap)
	avgFactor = float64(trigger)
}
