package sample_data

import "github.com/psinthorn/gostore/pb"

// return new keyboard object
func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  randomKeyboardlayout(),
		Backlit: randomBool(),
	}
	return keyboard
}
