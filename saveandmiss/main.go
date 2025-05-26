package main

import (
	"fmt"
)

func main() {
	fmt.Println(SaveAndMiss("123456789", 3))
	fmt.Println(SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Println(SaveAndMiss("", 3))
	fmt.Println(SaveAndMiss("hello you all ! ", 0))
	fmt.Println(SaveAndMiss("what is your name?", 0))
	fmt.Println(SaveAndMiss("go Exercise Save and Miss", -5))
}


func SaveAndMiss(arg string, num int) string {
	if num <= 0 {
		return arg
	}
	res := ""
	for i := 0; i < len(arg); i += num {
		if i+num <= len(arg) {
			res += arg[i : i+num]
		}else{
			res += arg[i:]
		}
		i += num
	}
	return res
}
