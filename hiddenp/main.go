package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	s1 := os.Args[1]
	s2 := os.Args[2]

	if s1 == "" || isHiden(s1, s2) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}

func isHiden(s1, s2 string) bool {
	indexs := []int{}
	isSorted := true
	for _, s := range s1 {
		one := true
		for j, ss := range s2 {
			if s == ss && one {
				one = false
				indexs = append(indexs, j)
			}
		}
	}
	fmt.Println(indexs)
	for i := range len(indexs)-1 {
		// fmt.Println(indexs[i] > indexs[i+1])
		if indexs[i] > indexs[i+1] {
			return false
		}
	}
	return isSorted
}
