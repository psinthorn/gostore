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

func randomKeyboardlayout() pb.Keyboard_Layout {

	switch rand.Intn(4) {
	case 0:
		return pb.Keyboard_UNKNOWN
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	case 3:
		return pb.Keyboard_AZERTY
	default:
		return pb.Keyboard_UNKNOWN
	}
}

func randomBool() bool {

	// ฟังค์ชั่นนี้จะใช้ rand.intn(2) ทำการสุ่มตัวเลขระหว่าง 0 และ 1 โดยจะทำการเปรียบเทียบกับตัวเลขด้านหลัง โดยที่ค่าที่คืนจะเป็น true หรือ false
	// 1. หากสุ่มตัวเลขได้ 1 เมื่อเปรียบเทียบกับเลข 1 ที่ด้านหลังก็จะเป็น 1 == 1 สมการก็เป็นจริงก็จะคืนค่าเป็น true
	// 2.  หากสุ่มตัวเลขได้ 0 เมื่อเปรียบเทียบกับเลข 1 ที่ด้านหลังก็จะเป็น 0 == 1 สมการไม่เป็นจริงก็จะคืนค่าเป็น false
	return rand.Intn(2) == 1
}
