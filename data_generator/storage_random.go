package sample_data

import (
	"math/rand"
	"time"

	"github.com/psinthorn/gostore/pb"
)

// by default random will use fix seed to create then some of random value will be remain the same
// we can fix it by tell random to use diffenrent seed to run
// by using below code
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomStorage() pb.Storage_Driver {

	n := rand.Intn(2)
	if n == 1 {
		return pb.Storage_SSD
	}
	return pb.Storage_HDD

}
