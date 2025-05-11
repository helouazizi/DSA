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

	n := len(input)
	visited := make([]bool, n)
	queue := []int{0}
	visited[0] = true

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		if int(input[pos])+pos == n-1 {
			return true
		}

		next := pos + int(input[pos])
		if next < n && !visited[next] {
			visited[next] = true
			queue = append(queue, next)
		}
	}

	return false

}
