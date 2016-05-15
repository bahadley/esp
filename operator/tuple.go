package operator

import (
	"encoding/json"
	"time"

	"github.com/bahadley/esp/log"
)

type (
	SensorTuple struct {
		Sensor    string    `json:"sensor"`
		Type      string    `json:"type"`
		Timestamp time.Time `json:"ts"`
		Data      float64   `json:"data"`
	}
)

func Unmarshal(msg string, st *SensorTuple) error {
	err := json.Unmarshal([]byte(msg), &st)
	if err != nil {
		log.Warning.Println(err.Error())
	}
	return err
}

func Marshal(sensor string, data float64) ([]byte, error) {
	var st SensorTuple

	st.Sensor = sensor
	st.Type = "TA"
	st.Timestamp = time.Now().UTC()
	st.Data = data

	msg, err := json.Marshal(st)
	if err != nil {
		log.Warning.Println(err.Error())
	}
	return msg, err
}
