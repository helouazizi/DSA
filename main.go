// main.go
package main

import "fmt"

func main() {
	array := make([]int, 1)
	/*n := len(array)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}*/
	mapp := make(map[string]int)
	fmt.Println(len(mapp))
	fmt.Println(cap(array))
	fmt.Println(len(array))
	fmt.Println(array)
	array = append(array, 6, 5, 6, 6, 6, 6)
	fmt.Println(cap(array))
	fmt.Println(len(array))

	fmt.Println(array)
	p := new([]int)
	fmt.Println(cap(*p))
	fmt.Println(len(*p))
	*p = append(*p, 5,5)
	fmt.Println(*p)
	fmt.Println(cap(*p))
	fmt.Println(len(*p))

}
