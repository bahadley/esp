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
	avg := 0.0

	inserted := insert(newTuple)
	if !inserted {
		log.Warning.Printf("Sensor %s tuple %d not inserted",
			newTuple.Sensor, newTuple.Timestamp)
	}

	if newTuple.Type == flushType {
		avg, agg = flush()
	} else {
		if inserted && window[bufSz-1] != nil {
			// Tuple was inserted and the window is full.
			avg = aggregate()
			agg = true
		}
	}

	if agg {
		aggTuple, err := Marshal(newTuple.Sensor, avg)
		if err != nil {
			log.Warning.Printf("Failed to marshal aggregate tuple for sensor: %s",
				newTuple.Sensor)
		} else {
			EgressChan <- aggTuple
		}
	}

	return nil
}

func insert(tmp *SensorTuple) bool {
	inserted := false

	for idx, st := range window {
		if inserted ||
			(!inserted && st != nil && tmp.Timestamp > st.Timestamp) {
			// Insert the new tuple and shift the subsequent tuples towards
			// the back of the window.  The last tuple will fall off if the
			// window is full.
			window[idx] = tmp
			tmp = st
			inserted = true
		} else if !inserted && st == nil {
			// Window is currently empty and this is the first arriving tuple, or ...
			// Out of order arrival and there is room at the back of the window.
			window[idx] = tmp
			inserted = true
			break
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

func flush() (float64, bool) {
	// Calculate the aggregation and flush all the tuples.
	// Take the aggregation from the front of the window in a flush.
	agg := true
	sum := 0.0

	for idx := uint32(0); idx < bufSz; idx++ {
		if idx < aggSz && window[idx] == nil {
			agg = false
		} else if idx < aggSz {
			sum += window[idx].Data
		}
		window[idx] = nil
	}

	return RoundDecimal(sum/float64(aggSz), 2), agg
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
