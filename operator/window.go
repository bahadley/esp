package operator

var (
	window []SensorTuple
	count  int
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
	if count == 2 {
		count = 0
		return (window[0].Data + window[1].Data) / 2.0, true
	}
	return 0.0, false
}

func init() {
	window = make([]SensorTuple, 4)
}
