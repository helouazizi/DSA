package main

import "github.com/01-edu/z01"

func main() {
	test := [10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'}
	PrintMemory(test)
	// z01.PrintRune('h')
}

func PrintMemory(text [10]byte) {
	hex := "0123456789abcdef"
	for i, char := range text {
		z01.PrintRune(rune(hex[char>>4])) // High nibble (first hex digit)

		z01.PrintRune(rune(hex[char&0x0F])) // Low nibble (second hex digit)
		z01.PrintRune(' ')
		if (i+1)%4 == 0 {
			z01.PrintRune('\n')
		}
	}
	z01.PrintRune('\n')
	for _, v := range text {
		if v <= 127 && v >= 32 {
			z01.PrintRune(rune(v))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}
