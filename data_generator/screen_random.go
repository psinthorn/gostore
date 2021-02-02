package sample_data

import (
	"math/rand"

	"github.com/psinthorn/gostore/pb"
)

func randomPanel() pb.Screen_Panel {
	switch rand.Intn(3) {
	case 1:
		return pb.Screen_IPS
	case 2:
		return pb.Screen_OLED
	default:
		return pb.Screen_UNKNOWN
	}
}

func randomResolution() pb.Screen_Resolution {
	return pb.Screen_Resolution{}
}
