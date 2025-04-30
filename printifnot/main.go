package main

import (
	"fmt"
)

func main() {
	fmt.Print(PrintIfNot("abcdefz"))
	fmt.Print(PrintIfNot("abc"))
	fmt.Print(PrintIfNot(""))
	fmt.Print(PrintIfNot("14"))
}


func PrintIfNot(str string) string {
	if str == ""{
		return "G\n"
	}

	if len(str)<3 {
		return "G\n"
	}

	return "Invalid Input\n"

}