package sample_data

import (
	"fmt"

	"github.com/golang/protobuf/ptypes"
	"github.com/psinthorn/gostore/pb"
)

// Return new laptop object
func NewLaptop() *pb.Laptop {
	fmt.Println("Generate New Laptop")
	newLaptop := &pb.Laptop{
		Id:          randomID(),
		Brand:       randomLaptopBrand(),
		Model:       randomLaptopModel(),
		Cpu:         NewCPU(),
		Memory:      NewMemory(),
		Gpus:        []*pb.GPU{NewGPU()},
		Storages:    []*pb.Storage{NewStorage()},
		Screen:      NewCreen(),
		Keyboard:    NewKeyboard(),
		Weight:      &pb.Laptop_WeightKg{WeightKg: randomFloat64(1.0, 3.0)},
		PriceUsd:    float64(randomInt(1500, 3000)),
		ReleaseYear: uint32(randomInt(2015, 2021)),
		CreatedAt:   ptypes.TimestampNow(),
		UpdatedAt:   ptypes.TimestampNow(),
	}
	return newLaptop
}
