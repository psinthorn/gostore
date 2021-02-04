package sample_data

import "github.com/psinthorn/gostore/pb"

// return new memory object
func NewMemory() *pb.Memory {
	memory := &pb.Memory{
		Value: uint64(randomInt(4, 32)),
		Unit:  pb.Memory_GIGABYTE,
	}
	return memory
}
