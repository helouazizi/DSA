package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println()
		return
	}
	str1 := os.Args[1]
	str2 := os.Args[2]
	seen := make(map[rune]bool)
	res := ""
	for _, v := range str1 + str2 {
		if !seen[v] {
			seen[v] = true
			res += string(v)
		}
	}
	fmt.Println(res)
}
