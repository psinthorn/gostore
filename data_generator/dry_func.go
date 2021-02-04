package sample_data

import (
	"math/rand"

	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
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

func randomID() string {
	return uuid.New
}

func randomLaptopBrand() string {
	return randomStringFromSet("Apple", "Dell", "HP", "IBM")
}

func randomLaptopModel() string {
	brand := randomLaptopBrand()
	switch brand {
	case "Apple":
		return randomStringFromSet("Macbook Pro", "Macbook Air", "Macbook")
	case "Dell":
		return randomStringFromSet("XPS", "Latitude", "Precision 3560")
	case "HP":
		return randomStringFromSet("ProBook 635 Aero G7", "Pavilion", "Spectre x360")
	default:
		return "UNKNOWN"
	}
}
