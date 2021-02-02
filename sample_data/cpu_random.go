package sample_data

import (
	"fmt"
	"math/rand"
)

func randomCpuBrand() string {
	return randomStringFromSet("Intel", "AMD", "Apple A13")
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	fmt.Println(n)
	if n == 0 {
		return ""
	}
	a1 := a[1]
	a2 := a[2]
	fmt.Println(a1)
	fmt.Println(a2)
	return a[rand.Intn(n)]
}
