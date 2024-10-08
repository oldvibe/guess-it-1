package cc

import "math"

func Variance(arr []float64) float64 {
	avg := Average(arr)
	sum := 0.0
	for _, value := range arr {
		sum += math.Pow(value-avg, 2)
	}
	return sum / float64(len(arr))
}
