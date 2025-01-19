// main.go
package main

import (
	"fmt"
	"math"
)

func main() {
	// this is the math.sqrt calculation
	fmt.Println(math.Sqrt(25), "default sqrt")
	fmt.Println(int(sqrt(25)), "our sqrt")
}

func sqrt(x float64) float64 {
	// we are using here the herons methods for calculating the squere root of a number
	guess := x / 2 // this the initial gesss
	// this formula need a tplerance
	tolerance := 0.1
	for {
		betterguess := (guess + (x / guess)) / 2.0
		if ads(betterguess-guess) < tolerance {
			return betterguess
		}
		guess = betterguess
	}
}

func ads(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
