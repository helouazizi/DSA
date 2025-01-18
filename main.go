// main.go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(IsPrime(-25))
}

func IsPrime(n int) bool {
	fmt.Println((math.Sqrt(float64(n))))
	if n <= 1 {
		return false
	}
	return true
}
