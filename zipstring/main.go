package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}

func ZipString(s string) string {
	res := ""
	n := len(s)
	if n == 0 {
		return s
	}

	for i:= 0 ; i < n ; {
		count := 1
		char:= s[i]
		for j := i+1; j < n && s[j] == char ; j++ {
			count++
		}
		res += strconv.Itoa(count)+string(char)
		i+=count
	}
	return res
}


