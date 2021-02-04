package sample_data

import (
	"math/rand"

	"github.com/psinthorn/gostore/pb"
)

func randomScreenPanel() pb.Screen_Panel {
	switch rand.Intn(3) {
	case 1:
		return pb.Screen_IPS
	case 2:
		return pb.Screen_OLED
	default:
		return pb.Screen_UNKNOWN
	}
}

func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9
	resolution := &pb.Screen_Resolution{
		Width:  uint32(width),
		Height: uint32(height),
	}

	return resolution
}
