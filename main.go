package main

import "fmt"

func main() {
	array := [5]int{9, 3, 6, 4, 5}

	n := len(array)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	fmt.Println(array)

}
