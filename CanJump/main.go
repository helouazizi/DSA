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

// lest code

func CanJump(input []uint) bool {
	if len(input) == 0 {
		return false
	}
	if len(input) == 1 {
		return true
	}

	n:= len(input)
	// Create a dp array to store the maximum reachable index at each position


	return false

}
