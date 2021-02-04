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
