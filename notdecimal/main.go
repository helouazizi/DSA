package main

import (
	"fmt"
)

func main() {
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
}
func NotDecimal(dec string) string {
	if dec == "" {
		return "\n"
	}

	hasDot := false
	valid := true
	dotIndex := -1
	first := true

	for i := range dec {
		c := dec[i]
		if c == '.' && first {
			hasDot = true
			dotIndex = i
			first = false
		}
		if !(c >= '0' && c <= '9') && c != '.' && c != '-' {
			valid = false
			break
		}
	}

	if !hasDot {
		return dec + "\n" // need to remove 000
	}

	if !valid {
		return dec + "\n"
	}

	beforDot := ""
	afterDot := ""

	for i := range dotIndex {
		beforDot += string(dec[i])
	}
	for i := dotIndex + 1; i < len(dec); i++ {
		afterDot += string(dec[i])
	}

	// chek only zeros
	onlyzero := true
	for i := range afterDot {
		if afterDot[i] != '0' {
			onlyzero = false
			break
		}
	}

	if onlyzero {
		return removeLeadingZeros(beforDot) + "\n"
	}

	// fmt.Println("befor", beforDot, "after", afterDot)
	return removeLeadingZeros(beforDot+afterDot) +"\n"

}

// func NotDecimal(dec string) string {
// 	if dec == "" {
// 		return "\n"
// 	}

// 	hasDot := false
// 	valid := true
// 	for i := 0; i < len(dec); i++ {
// 		c := dec[i]
// 		if !(c >= '0' && c <= '9') && c != '.' && c != '-' {
// 			valid = false
// 			break
// 		}
// 		if c == '.' {
// 			hasDot = true
// 		}
// 	}

// 	if !valid {
// 		return dec + "\n"
// 	}

// 	if !hasDot {
// 		return removeLeadingZeros(dec) + "\n"
// 	}

// 	dotIndex := -1
// 	for i := 0; i < len(dec); i++ {
// 		if dec[i] == '.' {
// 			dotIndex = i
// 			break
// 		}
// 	}

// 	beforeDot := ""
// 	for i := 0; i < dotIndex; i++ {
// 		beforeDot += string(dec[i])
// 	}

// 	afterDot := ""
// 	for i := dotIndex + 1; i < len(dec); i++ {
// 		afterDot += string(dec[i])
// 	}

// 	onlyZeros := true
// 	for i := 0; i < len(afterDot); i++ {
// 		if afterDot[i] != '0' {
// 			onlyZeros = false
// 			break
// 		}
// 	}

// 	if onlyZeros {
// 		return removeLeadingZeros(beforeDot) + "\n"
// 	}

// 	return removeLeadingZeros(beforeDot+afterDot) + "\n"
// }

func removeLeadingZeros(s string) string {
	isNegative := false
	start := 0

	if len(s) > 0 && s[0] == '-' {
		isNegative = true
		start = 1
	}

	// Skip leading zeros
	for start < len(s) && s[start] == '0' {
		start++
	}

	// If all were zeros, return "0"
	if start == len(s) {
		return "0"
	}

	result := ""
	if isNegative {
		result = "-"
	}
	for i := start; i < len(s); i++ {
		result += string(s[i])
	}
	return result
}
