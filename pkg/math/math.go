package math

import "math"

func sqrt32(n float32) float32 {
	return float32(math.Sqrt(float64(n)))
}

func pow32(n1, n2 float32, y float64) float32 {
	return float32(math.Pow(float64(n1-n2), y))
}
