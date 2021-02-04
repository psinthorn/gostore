package sample_data

import "github.com/psinthorn/gostore/pb"

// return new CPU object
func NewCPU() *pb.CPU {
	brand := randomCPUBrand()
	model := randomCPUModel(brand)
	numberOfCores := randomInt(2, 8)
	numberOfThread := randomInt(numberOfCores, 8)
	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5.0)

	cpu := &pb.CPU{
		Brand:         brand,
		Model:         model,
		NumberCores:   uint32(numberOfCores),
		NumberThreads: uint32(numberOfThread),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}
	return cpu
}

// return new GPU object
func NewGPU() *pb.GPU {
	brand := randomCPUBrand()
	model := randomCPUModel(brand)
	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 2.0)
	memory := &pb.Memory{
		Value: uint64(randomInt(2, 6)),
		Unit:  pb.Memory_GIGABYTE,
	}

	gpu := &pb.GPU{
		Brand:  randomGPUBrand(),
		Model:  model,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}

	return gpu

}
