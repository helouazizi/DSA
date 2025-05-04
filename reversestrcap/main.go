package main

import (
	"fmt"
	"os"
)

func main() {
	// if len(os.Args) != 3 {
	// 	fmt.Println()
	// 	return
	// }
	arg := os.Args[1:] 
	str := ""
	for i, s := range arg {
		str += s
		if i < len(arg)-1{
			str += "\n"
		}
	}
	// fmt.Println(str)
	arr := []string{}
	test := ""
	for _, v := range str {
		if v == ' ' {
			arr = append(arr, test)
			test = ""
		}
		test += string(v)
	}
	arr =append(arr, test)
	res := ""
	for _, v := range arr {
		// if v == "~"{
		// 	res += string('\n')
		// 	continue
		// }
		rest := v[:len(v)-1]
		last := v[len(v)-1:]
		res += tolower(string(rest))+touper(string(last))
		// fmt.Println(v,len(v),"hhhh")
	}
	
	// fmt.Println(arr)
	fmt.Println(res)
	// fmt.Println(touper(string(str[4])))

}

func tolower(s string) string {
	if s == ""{
		return s
	}
	res:= ""
	for _, v := range s {
		if v <= 'Z' && v >= 'A' {
			 res += string(v+('a'-'A'))
		}else {
			res += string(v)
		}
	}
	return res
}

func touper(s string) string {
	if s == ""{
		return s
	}
	res:= ""
	for _, v := range s {
		if v <= 'z' && v >= 'a' {
			 res += string(v+('A'-'a'))
		}else {
			res += string(v)
		}
	}
	return res
}
