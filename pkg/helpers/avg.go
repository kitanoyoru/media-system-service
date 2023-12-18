package helpers

func Avg(values []float64) float64 {
	var sum float64
	for _, v := range values {
		sum += v
	}

	return sum / float64(len(values))
}
