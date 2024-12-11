// test/main.go
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Room represents a room in the colony.
type Room struct {
	Name        string
	X, Y        int
	Connections []string
	IsStart     bool
	IsEnd       bool
}

// Colony represents the entire ant farm.
type Colony struct {
	Rooms map[string]*Room
	Start *Room
	End   *Room
	Ants  int
}

// AntMove represents a single ant's movement.
type AntMove struct {
	Ant   int
	Room  string
	Steps []string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]
	colony, err := ParseFile(filename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	paths := colony.BFS()
	if len(paths) == 0 {
		fmt.Println("Error: No paths found")
		os.Exit(1)
	}

	colony.DistributeAnts(paths)
}

// ParseFile parses the input file and constructs a Colony.
func ParseFile(filename string) (*Colony, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	colony := &Colony{Rooms: make(map[string]*Room)}
	reader := bufio.NewReader(file)
	var phase string

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			/*if errors.Is(err, os.EOF) {
				break
			}*/
			fmt.Println("error reading file: %w", err)
		}
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			if strings.HasPrefix(line, "##") {
				phase = line
			}
			continue
		}

		if colony.Ants == 0 {
			ants, err := strconv.Atoi(line)
			if err != nil || ants <= 0 {
				return nil, errors.New("invalid number of ants")
			}
			colony.Ants = ants
			continue
		}

		parts := strings.Fields(line)
		if len(parts) == 3 {
			// Room definition
			name := parts[0]
			if strings.HasPrefix(name, "L") || strings.HasPrefix(name, "#") {
				return nil, errors.New("invalid room name")
			}
			x, err1 := strconv.Atoi(parts[1])
			y, err2 := strconv.Atoi(parts[2])
			if err1 != nil || err2 != nil {
				return nil, errors.New("invalid room coordinates")
			}

			room := &Room{Name: name, X: x, Y: y}
			colony.Rooms[name] = room
			if phase == "##start" {
				room.IsStart = true
				colony.Start = room
			} else if phase == "##end" {
				room.IsEnd = true
				colony.End = room
			}
			phase = ""
		} else if len(parts) == 1 && strings.Contains(parts[0], "-") {
			// Link definition
			link := strings.Split(parts[0], "-")
			if len(link) != 2 {
				return nil, errors.New("invalid link format")
			}
			room1, room2 := link[0], link[1]
			if colony.Rooms[room1] == nil || colony.Rooms[room2] == nil {
				return nil, errors.New("link references unknown room")
			}
			colony.Rooms[room1].Connections = append(colony.Rooms[room1].Connections, room2)
			colony.Rooms[room2].Connections = append(colony.Rooms[room2].Connections, room1)
		} else {
			return nil, errors.New("invalid data format")
		}
	}

	if colony.Start == nil || colony.End == nil {
		return nil, errors.New("start or end room missing")
	}

	return colony, nil
}

// BFS finds the shortest paths from start to end.
func (c *Colony) BFS() [][]string {
	var paths [][]string
	queue := [][]string{{c.Start.Name}}
	visited := map[string]bool{}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		room := path[len(path)-1]

		if room == c.End.Name {
			paths = append(paths, path)
			continue
		}

		visited[room] = true
		for _, neighbor := range c.Rooms[room].Connections {
			if !visited[neighbor] {
				newPath := append([]string{}, path...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}

	return paths
}

// DistributeAnts distributes ants across paths and simulates their movements.
func (c *Colony) DistributeAnts(paths [][]string) {
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})

	pathLoad := make([]int, len(paths))
	assignments := map[int][]AntMove{}

	for ant := 1; ant <= c.Ants; ant++ {
		minLoadIndex := 0
		for i, load := range pathLoad {
			if load+len(paths[i]) < pathLoad[minLoadIndex]+len(paths[minLoadIndex]) {
				minLoadIndex = i
			}
		}
		pathLoad[minLoadIndex]++
		assignments[minLoadIndex] = append(assignments[minLoadIndex], AntMove{
			Ant:   ant,
			Room:  paths[minLoadIndex][0],
			Steps: paths[minLoadIndex],
		})
	}

	// Simulate turns
	turn := 0
	for {
		active := false
		fmt.Printf("Turn %d:\n", turn+1)
		for i, moves := range assignments {
			for j := range moves {
				if len(moves[j].Steps) > 1 {
					moves[j].Steps = moves[j].Steps[1:]
					fmt.Printf("L%d-%s ", moves[j].Ant, moves[j].Steps[0])
					active = true
				}
			}
			assignments[i] = moves
		}
		fmt.Println()
		if !active {
			break
		}
		turn++
	}
}
