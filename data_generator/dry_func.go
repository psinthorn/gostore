package sample_data

import (
	"math/rand"
)

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return "UNKNOWN"
	}
	return a[rand.Intn(n)]
}

func randomInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}
