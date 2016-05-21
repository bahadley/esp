package operator

import (
	"encoding/json"
	"time"

	"github.com/bahadley/esp/log"
)

const (
	aggregateId = "TA"
)

type (
	SensorTuple struct {
		Sensor    string  `json:"sensor"`
		Type      string  `json:"type"`
		Timestamp int64   `json:"ts"`
		Data      float64 `json:"data"`
	}
)

func Unmarshal(msg []byte, st *SensorTuple) error {
	err := json.Unmarshal(msg, &st)
	if err != nil {
		log.Warning.Println(err.Error())
	}
	return err
}

func Marshal(sensor string, data float64) ([]byte, error) {
	var st SensorTuple

	st.Sensor = sensor
	st.Type = aggregateId
	st.Timestamp = time.Now().UnixNano()
	st.Data = data

	msg, err := json.Marshal(st)
	if err != nil {
		log.Warning.Println(err.Error())
	}
	return msg, err
}
