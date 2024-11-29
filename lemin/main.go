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

// lets defin a struct to hold our room {node}
type Room struct {
	Name string
	X    int
	Y    int
}

// lets define a struct to hold our path {node - links}
// our graph sructure
type Graph struct {
	Rooms map[string]*Room
	Links map[string]*list.List
}

// so lets create a method to add a room to the graph
func (g *Graph) AddRoom(name string, x int, y int) {
	g.Rooms[name] = &Room{Name: name, X: x, Y: y}
}

// now lets add links method to add links

func (g *Graph) AddLink(from string, to string) {
	// lets check here if the rooms exist then
	// add the link
	if g.Links[from] == nil {
		g.Links[from] = list.New()
	}
	// add the link to the list
	g.Links[from].PushBack(to)
}

// / lets create our BFS function to find the shortest path
func (g *Graph) BFS(start string, end string) []string {

	// in first lets create a map to hold our visited rooms
	// avoiding infinite loops because the bfs is not fimilare with cycles graphs
	visited_rooms := make(map[string]bool)
	result := []string{}
	// lets create a queue to hold our nodes to visit
	nodes_queue := list.New()

	// lets add the start room to the queue
	nodes_queue.PushBack(start)
	// mark this room as visited
	visited_rooms[start] = true

	// know lets loop on the queue until  it is empty
	for nodes_queue.Len() > 0 {
		// lets get the first node from the queue
		current_room := nodes_queue.Front().Value.(string)
		// lets remove the node from the queue
		nodes_queue.Remove(nodes_queue.Front())
		// lets add the current room to the result
		result = append(result, current_room)

		// lets get the links of the current room
		for e := g.Links[current_room].Front(); e != nil; e = e.Next() {
			if !visited_rooms[e.Value.(string)] {
				// lets add the link to the queue
				nodes_queue.PushFront(e.Value.(string))
				// mark this room as visited
				visited_rooms[e.Value.(string)] = true

			}

		}
	}

	return result

}

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
	var numAnts int
	fmt.Scanf("%d\n", &numAnts)

	// Read rooms
	startRoom := &Room{}
	endRoom := &Room{}

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "L") {
			continue
		}

		if line == "##start" {
			startRoom = &Room{}
			continue
		}

		parts := strings.Fields(line)
		if len(parts) == 3 {
			x, _ := strconv.Atoi(parts[1])
			y, _ := strconv.Atoi(parts[2])
			room := &Room{Name: parts[0], X: x, Y: y}
			graph.AddRoom(room.Name, room.X, room.Y)

			if room == startRoom {
				startRoom = room
			}
		} else if line == "##end" {
			endRoom = &Room{}
			continue
		}
	}

	// Read links
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "L") {
			continue
		}
		parts := strings.Split(line, "-")
		if len(parts) == 2 {
			graph.AddLink(parts[0], parts[1])
		}
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
		fmt.Printf("ERROR: invalid data format\n")
		os.Exit(1)
	}

	startRoom := graph.Rooms["##start"]
	endRoom := graph.Rooms["##end"]

	if startRoom == nil || endRoom == nil {
		fmt.Println("ERROR: no start room found")
		os.Exit(1)
	}

	path := BFS(graph, startRoom.Name, endRoom.Name)

	fmt.Printf("%d\n", len(path))
	for _, room := range path {
		fmt.Printf("%s ", graph.Rooms[room])
	}
	fmt.Println()

	// Print links
	for from, links := range graph.Links {
		fmt.Printf("%s-", from)
		for i := 0; i < links.Len(); i++ {
			fmt.Printf(" %s", links.Element(i).(string))
		}
		fmt.Println()
	}
}
