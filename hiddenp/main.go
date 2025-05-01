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

	isHiden(s1, s2)
}

func isHiden(s1, s2 string)  {
	if s1 == "" {
		fmt.Println(1)
	}

	j := 0
	for i := 0; i < len(s2) && j < len(s1); i++ {
		if s1[j] == s2[i] {
			j++
		}
	}
	if j == len(s1){
		fmt.Println(1)
	}else{
		fmt.Println(0)
	}
}
