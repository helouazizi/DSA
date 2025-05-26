package main

import (
	"fmt"
)

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2))

	input3 := []uint{0}
	fmt.Println(CanJump(input3))
}

func CanJump(input []uint) bool {
	if len(input) == 0 {
		return false
	}
	if len(input) == 1 {
		return true
	}

	pos := 0
	for pos < len(input)-1 {
		step := int(input[pos])
		if pos+step >= len(input) || step == 0 {
			return false
		}
		pos += step
	}
	return pos == len(input)-1
}

// if len(input) == 0 {
// 	return false
// }
// pos := 0
// for pos < len(input)-1 {
// 	step := int(input[pos])
// 	if step == 0 || pos+step >= len(input) {
// 		return false
// 	}
// 	pos += step
// }
// return pos == len(input)-1
