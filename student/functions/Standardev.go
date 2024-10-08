package cc

import "math"

func StandardDev(arr []float64) float64 {
	return math.Sqrt(Variance(arr))
}
