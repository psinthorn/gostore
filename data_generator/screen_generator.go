package sample_data

import (
	"github.com/psinthorn/gostore/pb"
)

// return new creen object
func NewCreen() *pb.Screen {
	screen := &pb.Screen{
		SizeInch:   float32(randomInt(13, 32)),
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}
	return screen
}
