package src

func Mean(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}

// Variance calculates the variance of a slice of float64 values.
func Variance(data []float64) float64 {
	mean := Mean(data)
	var sumSquares float64
	for _, value := range data {
		sumSquares += (value - mean) * (value - mean)
	}
	return sumSquares / float64(len(data))
}
