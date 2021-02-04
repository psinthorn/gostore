package sample_data

import "github.com/psinthorn/gostore/pb"

// return new storage object
func NewStorage() *pb.Storage {
	storage := &pb.Storage{

		Driver: randomStorage(),
		Memory: &pb.Memory{
			Value: uint64(randomInt(2, 6)),
			Unit:  pb.Memory_GIGABYTE,
		},
	}

	return storage
}
