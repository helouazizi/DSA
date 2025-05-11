package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Println(FifthAndSkip("This is a short sentence"))
	fmt.Println(FifthAndSkip("1234"))
}

func FifthAndSkip(str string) string {
	if len(str) == 0 {
		return "\n"
	}

	nospace := strings.ReplaceAll(str, " ", "")
	if len(nospace) < 5 {
		return "Invalid Input"
	}

	result := ""
	count := 0
	temp := ""

	for _, v := range str {
		if v == ' ' {
			continue
		}

		if count == 5 {
			result += temp + " "
			temp = ""
			count = 0
			continue
		} 
		temp += string(v)
		count++

	}
	result += temp

	return result

}
