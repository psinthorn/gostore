package sample_data

import (
	"math/rand"
	"time"
)

// by default random will use fix seed to create then some of random value will be remain the same
// we can fix it by tell random to use diffenrent seed to run
// by using below code
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Random CPU Brand
func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD", "Apple A13")
}

// Random CPU model refering by Brand Name
func randomCPUModel(brand string) string {
	switch brand {
	case "Intel":
		return randomStringFromSet(
			"Xeon E-2286M",
			"Core i-9-9980HK",
			"Core i-7-9750H",
			"Core i-5-9400F",
			"Core i-3-1005G1",
		)
	case "AMD":
		return randomStringFromSet(
			"Ryzen 7 PRO 2700U",
			"Ryzen 5 PRO 3500U",
			"Ryzen 3 PRO 2300GE",
		)

	case "Apple":
		return randomStringFromSet(
			"A4 APL2298",
			"A5 APL7498",
			"A8 APL1011",
			"A13 Bionic APL1W01",
			"T1 APL1023",
			"T2 APL1027",
		)
	default:
		return randomStringFromSet(
			"Xeon x-xxxxx",
			"Core x-x-xxxxx",
			"Ryzen x xx xxxxx",
			"A13 xxxxxxx",
			"T1 xxxxxxx",
			"T2 xxxxxxx",
		)
	}
}

func randomGPUBrand() string {
	return randomStringFromSet("NVIDIA", "AMD")
}

func randomGPUModel(brand string) string {
	if brand == "NVIDIA" {
		return randomStringFromSet(
			"GeForce 8400M G",
			"GeForce GTX 1650",
			"GeForce MX330",
		)
	}

	return randomStringFromSet(
		"Radeon™ RX 480 Graphics",
		"Radeon™ 540 Graphics",
	)
}
