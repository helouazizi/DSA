package main

import "fmt"

func main() {
	num := 55
	num_as_text := Itoa(num)
	fmt.Printf("num:%d %T\n", num, num)
	fmt.Printf("num_as_text:%s %T\n", num_as_text, num_as_text)
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
