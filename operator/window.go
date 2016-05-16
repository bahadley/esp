package operator

import (
	"os"
	"strconv"

	"github.com/bahadley/esp/log"
)

const (
	defaultLength     = 4
	defaultTriggerQty = 2

	envWinLen  = "ESP_WINDOW_LENGTH"
	envWinTrig = "ESP_WINDOW_TRIGGER"
)

var (
	window  []*SensorTuple
	count   int
	trigger int
)

func WindowInsert(msg string) ([]byte, bool) {
	var rslt []byte
	var newTuple, tmp *SensorTuple

	newTuple = new(SensorTuple)
	err := Unmarshal(msg, newTuple)
	if err != nil {
		return rslt, false
	}

	for idx, st := range window {
		if idx == 0 {
			window[idx] = newTuple
		} else {
			window[idx] = tmp
		}
		tmp = st
	}

	count++
	if count == trigger {
		count = 0
		rslt, err = Marshal(newTuple.Sensor, average())
		return rslt, true
	}
	return rslt, false
}

func average() float64 {
	sum := 0.0
	for _, st := range window[0:trigger] {
		sum += st.Data
	}
	return sum / float64(trigger)
}

func init() {
	var err error
	var winlen int

	envVal := os.Getenv(envWinLen)
	if len(envVal) == 0 {
		winlen = defaultLength
	} else {
		winlen, err = strconv.Atoi(envVal)
		if err != nil {
			log.Error.Fatalf("Invalid environment variable: %s", envWinLen)
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

	window = make([]*SensorTuple, winlen)
}
