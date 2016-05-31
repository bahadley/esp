package operator

import (
	"github.com/bahadley/esp/log"
)

const flushType = "F"

var (
	// Invariant:  Tupls are in descending order by SensorTuple.Timestamp.
	window []*SensorTuple

	// Length of window.
	bufSz uint32
	// Number of tuples that will compose the aggregate calculation.
	aggSz uint32
)

func windowInsert(msg []byte) error {
	newTuple, err := Unmarshal(msg)
	if err != nil {
		log.Warning.Printf("Failed to unmarshal tuple: %s", msg)
		return err
	}

	agg := false

	if newTuple.Type == flushType {
		agg = true
	} else {
		if insert(newTuple) {
			agg = true
		} else {
			log.Warning.Printf("Sensor %s tuple %d not inserted",
				newTuple.Sensor, newTuple.Timestamp)
		}
	}

	if agg && window[bufSz-1] != nil {
		// Flushing or a tuple was inserted and the window is full.
		avg := aggregate()
		aggTuple, err := Marshal(newTuple.Sensor, avg)
		if err != nil {
			log.Warning.Printf("Failed to marshal aggregate tuple for sensor: %s",
				newTuple.Sensor)
		} else {
			EgressChan <- aggTuple
		}
	}

	if newTuple.Type == flushType {
		// Complete the flush of all tuples.
		flush()
	}

	return nil
}

func insert(tmp *SensorTuple) bool {
	inserted := false

	if window[0] == nil {
		// This case only occurs for the very first tuple received by the operator.
		// Handling this special case simplifies the remaining logic.
		window[0] = tmp
		inserted = true
	} else {
		for idx, st := range window {
			if inserted ||
				(!inserted && st != nil && tmp.Timestamp > st.Timestamp) {
				// Insert the new tuple and shift the subsequent tuples towards
				// the back of the window.  The last tuple will fall off if the
				// window is full.
				window[idx] = tmp
				tmp = st
				inserted = true
			}
		}
	}

	return inserted
}

func aggregate() float64 {
	// Calculate the aggregation and flush the tuples that were used.
	sum := 0.0

	for idx := bufSz - aggSz; idx < bufSz; idx++ {
		sum += window[idx].Data
		window[idx] = nil
	}

	return RoundDecimal(sum/float64(aggSz), 2)
}

func flush() {
	for i := uint32(0); i < bufSz; i++ {
		window[i] = nil
	}
}

func init() {
	bufSz = WindowSize()
	aggSz = AggregateSize()

	if bufSz < aggSz {
		log.Error.Fatalf("Invalid environment variables: %s < %s",
			envWinSz, envAggSz)
	}

	window = make([]*SensorTuple, bufSz)
}
