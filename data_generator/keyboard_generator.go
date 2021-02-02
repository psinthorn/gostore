package sample_data

import "github.com/psinthorn/gostore/pb"

// NewKeyboard() will returns new keyboard object
func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  randomKeyboardlayout(),
		Backlit: randomBacklitBoolean(),
	}
}
