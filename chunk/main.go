package main

import "fmt"


func main() {
		
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)

}
func Chunk(slice []int, size int) {
	if size == 0 {
		fmt.Println()
		return
	}
	if len(slice) == 0 {
		fmt.Println(slice)
		return
	}

	chunks := [][]int{}
	tst := []int{}
	for i:= range slice {
		
		if i % size == 0  && i > 0{
			chunks = append(chunks, tst)
			tst = []int{}
		}
		tst = append(tst, slice[i])
	}
	chunks = append(chunks, tst)

	fmt.Println(chunks) 

}