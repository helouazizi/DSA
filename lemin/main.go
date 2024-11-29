// lemin/main.go
package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Room struct to hold room information
type Room struct {
	Name string
	X    int
	Y    int
}

// Graph struct to hold rooms and links
type Graph struct {
	Rooms map[string]*Room
	Links map[string]*list.List
}

// AddRoom method to add a room to the graph
func (g *Graph) AddRoom(name string, x int, y int) {
	g.Rooms[name] = &Room{Name: name, X: x, Y: y}
}

// AddLink method to add a link between two rooms
func (g *Graph) AddLink(from string, to string) {
	if g.Links[from] == nil {
		g.Links[from] = list.New()
	}
	g.Links[from].PushBack(to)
}

// BFS method to find the path from start to end room
func (g *Graph) BFS(start string, end string) []string {
	visitedRooms := make(map[string]bool)
	result := []string{}
	nodesQueue := list.New()

	nodesQueue.PushBack(start)
	visitedRooms[start] = true

	for nodesQueue.Len() > 0 {
		currentRoom := nodesQueue.Front().Value.(string)
		nodesQueue.Remove(nodesQueue.Front())
		result = append(result, currentRoom)

		// Check if there are links for the current room
		if g.Links[currentRoom] == nil {
			continue // Skip if no links
		}

		for e := g.Links[currentRoom].Front(); e != nil; e = e.Next() {
			neighbor := e.Value.(string)
			if !visitedRooms[neighbor] {
				nodesQueue.PushBack(neighbor)
				visitedRooms[neighbor] = true
			}
		}
	}

	return result
}

// readInput function to read graph data from a file
func readInput(filename string) (*Graph, error) {
	graph := &Graph{
		Rooms: make(map[string]*Room),
		Links: make(map[string]*list.List),
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read number of ants
	/*var numAnts int
	if scanner.Scan() {
		numAnts, _ = strconv.Atoi(scanner.Text())
	}*/

	var startRoomName, endRoomName string

	// Read rooms
	for scanner.Scan() {
		line := scanner.Text()
		if line == "##start" {
			scanner.Scan()
			startRoomData := strings.Fields(scanner.Text())
			if len(startRoomData) == 3 {
				startRoomName = "start" // Assign a name for the start room
				x, _ := strconv.Atoi(startRoomData[0])
				y, _ := strconv.Atoi(startRoomData[1])
				graph.AddRoom(startRoomName, x, y)
			}
			continue
		} else if line == "##end" {
			scanner.Scan()
			endRoomData := strings.Fields(scanner.Text())
			if len(endRoomData) == 3 {
				endRoomName = "end" // Assign a name for the end room
				x, _ := strconv.Atoi(endRoomData[0])
				y, _ := strconv.Atoi(endRoomData[1])
				graph.AddRoom(endRoomName, x, y)
			}
			continue
		}

		// Read other rooms
		roomData := strings.Fields(line)
		if len(roomData) == 3 {
			roomName := roomData[0]
			x, _ := strconv.Atoi(roomData[1])
			y, _ := strconv.Atoi(roomData[2])
			graph.AddRoom(roomName, x, y)
			continue
		}

		// Read links
		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			if len(parts) == 2 {
				graph.AddLink(parts[0], parts[1])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return graph, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: lem-in <input_file>")
		os.Exit(1)
	}

	filename := os.Args[1]
	graph, err := readInput(filename)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}

	startRoom := graph.Rooms["start"]
	endRoom := graph.Rooms["end"]

	if startRoom == nil || endRoom == nil {
		fmt.Println("ERROR: no start or end room found")
		os.Exit(1)
	}

	fmt.Println("Starting BFS...")

	// Debugging: Print the rooms and links
	fmt.Println("Rooms:")
	for _, room := range graph.Rooms {
		fmt.Printf("Name: %s, Coordinates: (%d, %d)\n", room.Name, room.X, room.Y)
	}

	fmt.Println("Links:")
	for from, links := range graph.Links {
		fmt.Printf("%s-", from)
		for e := links.Front(); e != nil; e = e.Next() {
			fmt.Printf(" %s", e.Value.(string))
		}
		fmt.Println()
	}

	path := graph.BFS(startRoom.Name, endRoom.Name)

	fmt.Printf("Path length: %d\n", len(path))
	for _, room := range path {
		fmt.Printf("%s ", graph.Rooms[room].Name)
	}
	fmt.Println()
}
