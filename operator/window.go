package operator

import (
    "encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/bahadley/esp/log"
)

type (
	gMsg struct {
		Sensor     string     `json:"sensor"`
		Type       string     `json:"type"`
        Timestamp  time.Time  `json:"ts"`
        Data       float64    `json:"data"`
	}
)

func Window(ingest chan string) {
	for {
		msg := <-ingest
		var gm gMsg
        err := json.NewDecoder(strings.NewReader(msg)).Decode(&gm)
        if err != nil {
			log.Logoutput(log.ErrPrefix, err.Error())
			continue
        }
		log.Logoutput(log.InfoPrefix, 
			fmt.Sprintf("Windowed: %s,%s,%s,%f", gm.Sensor, 
				gm.Type, gm.Timestamp, gm.Data))
	}
}
