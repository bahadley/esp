package operator

import (
	"time"
)

type (
	SensorTuple struct {
		Sensor    string    `json:"sensor"`
		Type      string    `json:"type"`
		Timestamp time.Time `json:"ts"`
		Data      float64   `json:"data"`
	}
)
