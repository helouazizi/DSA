package main

import "fmt"

func main() {
	num := 55
	text := "111"
	num_as_text := Itoa(num)
	inttext := Atoi(text)
	fmt.Printf("num:%d %T\n", num, num)
	fmt.Printf("num_as_text:%s %T\n", num_as_text, num_as_text)
	fmt.Printf("inttext:%d %T\n", inttext, inttext)
}

func Itoa(num int) string {
	var num_as_text string
	for num > 0 {
		rem := num % 10
		num_as_text += string(rune(rem) + '0')
		num /= 10
	}
	return num_as_text
}

func Atoi(num string) int {
	intt := 0
	for _, v := range num {
		intt = intt*10 + (int(v) - '0')
	}
	return intt
}
