package operator

import (
	"os"
	"strconv"

	"github.com/bahadley/esp/log"
)

const (
	defaultLength  = 4
	defaultTrigger = 2

	envWinLen  = "ESP_WINDOW_LENGTH"
	envWinTrig = "ESP_WINDOW_TRIGGER"
)

var (
	window  []*SensorTuple
	count   int
	trigger int
)

func WindowAppend(msg string) error {
	var newTuple, tmp *SensorTuple

	newTuple = new(SensorTuple)
	err := Unmarshal(msg, newTuple)
	if err != nil {
		log.Warning.Printf("Failed to unmarshal tuple: %s", msg)
		return err
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
		rslt, err := Marshal(newTuple.Sensor, average())
		if err != nil {
			log.Warning.Printf("Failed to marshal aggregate tuple for sensor: %s",
				newTuple.Sensor)
			return err
		}
		EgressChan <- rslt
	}

	return nil
}

func average() float64 {
	sum := 0.0
	for i := 0; i < trigger; i++ {
		sum += window[i].Data
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
			log.Error.Fatalf("Invalid environment variable: %s",
				envWinLen)
		}
	}

	if winlen <= 0 {
		log.Error.Fatalf("Invalid environment variable: %s <= 0",
			envWinLen)
	}

	envVal = os.Getenv(envWinTrig)
	if len(envVal) == 0 {
		trigger = defaultTrigger
	} else {
		trigger, err = strconv.Atoi(envVal)
		if err != nil {
			log.Error.Fatalf("Invalid environment variable: %s",
				envWinTrig)
		}
	}

	if winlen < trigger {
		log.Error.Fatalf("Invalid environment variables: %s < %s",
			envWinLen, envWinTrig)
	}

	window = make([]*SensorTuple, winlen)
}
