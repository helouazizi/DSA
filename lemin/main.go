// lemin/main.go
// lemin/lem-in/main.go
// main.go
package main

import (
	"fmt"
	"os"

	"lem-in/helpers" // add somthig to test in my life
)

func main() {
	// lets check the argumments to protect the main function
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . exemple.txt")
		return
	}
	farm := &helpers.Farm{}

	// now lets read the file to extract ou data
	fileName := os.Args[1]
	err := farm.ReadFile(fileName)
	if err != nil {
		// fmt.Println("ERROR: invalid data format")
		fmt.Println(err)
		return
	}

	// fmt.Println(farm.Ants)

	//fmt.Println(farm.StartRoom)
	//fmt.Println(farm.EndRoom)
	// fmt.Println(farm.Rooms)

	paths := farm.Path_Finder()
	fmt.Println(paths)
	//fmt.Println(farm.Links)
	//farm.DistributeAnts(paths)
	fmt.Println("good data")
}
