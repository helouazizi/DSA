// main.go
package main

import (
	"fmt"
	"math"
)

func main() {
	// this is the math.sqrt calculation
	fmt.Println(int(math.Sqrt(25)), "default sqrt")
	// this is the herons sqrt
	fmt.Println(int(HeronsSqrt(25)), "herons sqrt")
	//this is the binary saerch sqrt
	fmt.Println(int(binarySearchSqrt(25)), "binary search sqrt")
}

func HeronsSqrt(x float64) float64 {
	// we are using here the herons methods for calculating the squere root of a number
	if x < 0 {
		return -1 // Return -1 for invalid input (negative numbers)
	}
	if x == 0 || x == 1 {
		return x
	}
	guess := x / 2 // this the initial gesss
	// this formula need a tplerance
	tolerance := 0.00001
	for {
		betterguess := (guess + (x / guess)) / 2.0
		if ads(betterguess-guess) < tolerance {
			return betterguess
		}
		guess = betterguess
	}
}

func binarySearchSqrt(x float64) float64 {
	// so the binary search consist of using
	// lower , higher and midle variables with tolerance number of course
	if x < 0 {
		return -1 // Return -1 for invalid input (negative numbers)
	}
	if x == 0 || x == 1 {
		return x
	}

	low, high := 0.0, x
	tolerance := 0.00001
	var mid float64
	for (high - low) > tolerance {
		mid = (high + low) / 2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			low = mid
		} else {
			high = mid
		}
	}

	return mid
}

func ads(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
