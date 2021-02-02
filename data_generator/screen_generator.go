package sample_data

import (
	"fmt"

	"github.com/psinthorn/gostore/pb"
)

func NewCreen() *pb.Screen {
	fmt.Println("Implement me please")

	return &pb.Screen{
		SizeInch:   float32(randomInt(15, 32)),
		Resolution: randomResolution(),
		Panel:      randomPanel(),
		Multitouch: randomBool(),
	}
}
