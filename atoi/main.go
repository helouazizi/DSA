package main

import "fmt"

func main() {
	num := "55"
	num_as_text := Atoi(num)
	fmt.Printf("num:%s %T\n", num, num)
	fmt.Printf("num_as_text:%d %T\n", num_as_text, num_as_text)
}

func Atoi(num string) int {
	intt := 0
	for _, v := range num {
		intt = intt*10 + (int(v) - '0')
	}
	return intt
}
