package main

import (
	"fmt"
)

func main() {
	fmt.Println(Gcd(42, 10))
	fmt.Println(Gcd(42, 12))
	fmt.Println(Gcd(14, 77))
	fmt.Println(Gcd(17, 3))
}

func Gcd(a, b uint) uint {
	if a == 0 || b == 0 {
		return 0
	}

	for b != 0 {
		a, b = b, a%b
	}
	return a

}
