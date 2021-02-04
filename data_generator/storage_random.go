package sample_data

import (
	"math/rand"

	"github.com/psinthorn/gostore/pb"
)

func randomStorage() pb.Storage_Driver {

	n := rand.Intn(2)
	if n == 1 {
		return pb.Storage_SSD
	}
	return pb.Storage_HDD

}
