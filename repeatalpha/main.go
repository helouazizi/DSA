package main

import (
	"fmt"
)

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}

func RepeatAlpha(s string) string {

	res := ""
	// start := rune('a')
	for i := range len(s) {
		if s[i] <= 'z' && s[i] >= 'a' {
			dif := rune(s[i]) - rune('a')
			for range int(dif + 1) {
				res += string(s[i])
			}
		} else if s[i] <= 'Z' && s[i] >= 'A' {
			dif := rune(s[i]) - rune('A')
			for range int(dif + 1) {
				res += string(s[i])
			}
		} else {
			res += string(s[i])
		}

	}

	// fmt.Println(rune('a'))
	return res
}
