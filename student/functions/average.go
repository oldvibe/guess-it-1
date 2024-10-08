package cc

func Average(arr []float64) float64 {
	sum := 0.0
	for _, value := range arr {
		sum += value
	}
	return sum / float64(len(arr))
}
