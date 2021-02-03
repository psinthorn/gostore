package sample_data

import "github.com/psinthorn/gostore/pb"

func NewMemory() *pb.Memory {
	memory := &pb.Memory{
		Value: uint64(randomInt(4, 16)),
		Unit:  pb.Memory_GIGABYTE,
	}
	return memory
}
