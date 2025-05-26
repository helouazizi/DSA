package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		return
	}
	text1 := args[1]
	text2 := args[2]
	// combin := text1 + text2
	res := ""
	for i := range text1 {
		if !Exist(res, string(text1[i])) && Exist(text2,string(text1[i])) {
			res += string(text1[i])
		}
	}
	fmt.Println(res)

}

func Exist(s, char string) bool {
	for i := range s {
		if string(s[i]) == char {
			return true
		}
	}
	return false
}
