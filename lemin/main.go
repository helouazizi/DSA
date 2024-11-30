// lemin/main.go
package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

// this struct represent a room
// with it name and coordonetes (x,y)
type Room struct {
	Name string
	X    int
	Y    int
}

// this struct reperesent an undirect graph
// data structure to holde rooms as a map
// and links between them as a map
type Graph struct {
	Rooms map[string]*Room
	Links map[string]*list.List
}

func main() {
	// lest check the argements
	if len(os.Args) != 2 {
		fmt.Println("[Usage]: ./lemin exemple.txt")
		os.Exit(1)
	}

	filename := os.Args[1]

	// lets read the file
	graph, numOfants, err := readfile(filename)
}

func readfile(filename string) (*Graph, int, error) {
	// lets open the file
	// and read it line by line
	// we will store the rooms and links in the graph struct
	// we will also count the number of ants
	file, err := os.Open(filename)
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// we will read the file line by line
	// the first line is the number of ants
	for scanner.Scan() {
		line := scanner.Text()
		
	}

}
