// test/main.go
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Room struct {
	X, Y string
}

type AntMove struct {
	Ant   int
	Room  string
	Steps []string
}

type Farm struct {
	Rooms              map[string]Room
	Links              map[string][]string
	StartRoom, EndRoom string
	Ants               int
	FileSize           int64
}

func (F *Farm) ReadFile(fileName string) error {
	fileinfo, err := os.Stat(fileName)
	if err != nil {
		return err
	}
	F.FileSize = fileinfo.Size()

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	if F.Rooms == nil {
		F.Rooms = make(map[string]Room)
	}
	if F.Links == nil {
		F.Links = make(map[string][]string)
	}

	scanner := bufio.NewScanner(file)
	state := "start"
	var roomDefinitions []string
	var linkDefinitions []string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "##") {
			continue
		}

		if strings.HasPrefix(line, "##") {
			if line == "##start" {
				state = "start-room"
			} else if line == "##end" {
				state = "end-room"
			}
			continue
		}

		parts := strings.Fields(line)
		if state == "start-room" || state == "end-room" {
			if len(parts) != 3 {
				return errors.New("invalid room definition")
			}
			roomName := parts[0]
			F.Rooms[roomName] = Room{X: parts[1], Y: parts[2]}
			if state == "start-room" {
				F.StartRoom = roomName
			} else {
				F.EndRoom = roomName
			}
			state = "normal"
		} else if len(parts) == 3 {
			roomDefinitions = append(roomDefinitions, line)
			roomName := parts[0]
			F.Rooms[roomName] = Room{X: parts[1], Y: parts[2]}
		} else if len(parts) == 1 && strings.Contains(parts[0], "-") {
			linkDefinitions = append(linkDefinitions, parts[0])
		}
	}

	if F.StartRoom == "" || F.EndRoom == "" {
		return errors.New("missing start or end room")
	}

	for _, link := range linkDefinitions {
		rooms := strings.Split(link, "-")
		if len(rooms) != 2 {
			return errors.New("invalid link format")
		}
		room1, room2 := rooms[0], rooms[1]
		if _, ok := F.Rooms[room1]; !ok {
			return fmt.Errorf("link to unknown room: %s", room1)
		}
		if _, ok := F.Rooms[room2]; !ok {
			return fmt.Errorf("link to unknown room: %s", room2)
		}
		F.Links[room1] = append(F.Links[room1], room2)
		F.Links[room2] = append(F.Links[room2], room1)
	}

	return nil
}

func (F *Farm) Path_Finder() [][]string {
	stack := [][]string{{F.StartRoom}}
	var result [][]string

	for len(stack) > 0 {
		path := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node := path[len(path)-1]

		if node == F.EndRoom {
			result = append(result, append([]string{}, path...))
			continue
		}

		for _, neighbor := range F.Links[node] {
			if !contains(path, neighbor) {
				newPath := append([]string{}, path...)
				newPath = append(newPath, neighbor)
				stack = append(stack, newPath)
			}
		}
	}

	return result
}

func contains(path []string, connection string) bool {
	for _, connected := range path {
		if connected == connection {
			return true
		}
	}
	return false
}

// lemin/helpers/helpers.go

// DistributeAnts distributes ants across paths without collisions and tracks their movements.
func DistributeAnts(ants int, paths [][]int) {
	// Step 1: Sort paths by length (shortest to longest)
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})

	// Step 2: Assign ants to paths
	antAssignments := make([][]int, len(paths))
	antQueue := make([]int, ants)
	for i := 0; i < ants; i++ {
		antQueue[i] = i + 1
	}

	// Assign ants to paths as evenly as possible
	for len(antQueue) > 0 {
		for i := range paths {
			if len(antQueue) == 0 {
				break
			}
			ant := antQueue[0]
			antQueue = antQueue[1:]
			antAssignments[i] = append(antAssignments[i], ant)
		}
	}

	// Step 3: Simulate movement
	fmt.Println("Ant movements:")
	active := true
	turn := 1
	antPositions := make(map[int]int) // Tracks the current position index of each ant

	// Initialize ant positions
	for i, _ := range paths {
		for _, ant := range antAssignments[i] {
			antPositions[ant] = 0 // Start at the beginning of the path
		}
	}

	// Simulate turns until all ants reach the end
	for active {
		active = false
		fmt.Printf("Turn %d:\n", turn)
		for i, path := range paths {
			for _, ant := range antAssignments[i] {
				pos := antPositions[ant]
				if pos < len(path)-1 { // If the ant has not reached the end
					active = true
					antPositions[ant]++
					fmt.Printf("Ant %d-%d ", ant, path[antPositions[ant]])
				}
			}
		}
		fmt.Println() // Newline after each turn
		turn++
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	filename := os.Args[1]
	
	graph, ants, err := ParseFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File parsed successfully.")

	// Find paths (ensure `FindPaths` works as expected)
	paths := FindPaths(graph)
	fmt.Println("Paths found:")
	for _, path := range paths {
		fmt.Println(path)
	}

	// Simulate ants distribution
	DistributeAnts(ants, paths)
}
