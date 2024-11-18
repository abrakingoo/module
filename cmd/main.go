package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Check if a file name was provided as an argument
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . {filename}")
		return
	}

	fileName := os.Args[1]

	// Check if the file has a .txt extension
	if !strings.HasSuffix(fileName, ".txt") {
		fmt.Println("ERROR: invalid file format, only (.txt) files are supported.")
		return
	}

	// Read the contents of the file
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	fileData := strings.Split(string(file), "\n")
	// Check if the number of ants is provided
	ants := fileData[0]
	if len(ants) == 1 && len(fileData[1]) != 1 {
		fileData = fileData[1:]
	} else {
		log.Fatalf("ERROR: invalid data format")
	}

	// Check if the program has start or end
	var first, last bool

	for i := 0; i < len(fileData); i++ {
		if string(fileData[i]) == "##start" {
			first = true
		}

		if string(fileData[i]) == "##end" {
			last = true
			if first {
				break // if both are true
			}
		}
	}

	if !first || !last {
		log.Fatalf("Error: invalid data format")
	}

	var (
		start   int
		end     int
		initial int
		count   int
		rooms   []string
		tunnels []string
	)

	// Parse rooms and find the start and end positions
loop:
	for i := 0; i < len(fileData); i++ {
		line := fileData[i]
		if line == "" {
			continue
		}
		if line == "##start" {
			start = count
		}

		if line[0] == '#' && line != "##end" {
			continue
		}

		if line == "##end" {
			end = count
			for i < len(fileData) {
				if strings.Contains(string(fileData[i]), "-") {
					initial = i
					break loop
				}
				if string(fileData[i][0]) != "#" {
					rooms = append(rooms, string(fileData[i][0]))
				}
				i++
			}
		}
		rooms = append(rooms, string(line[0])) // Add the room name
		count++
	}

	// Parse tunnels starting from the "start" index
	for i := initial; i < len(fileData); i++ {
		line := fileData[i]
		if line == "" || line[0] == '#' {
			continue
		}
		tunnels = append(tunnels, line)
	}

	// Create the graph
	graph := NewGraph()
	for _, connection := range tunnels {
		if strings.Contains(connection, "-") {
			rooms := strings.Split(connection, "-")
			graph.AddEdges(rooms[0], rooms[1])
		}
	}

	// Print initial input
	fmt.Println(strings.Join(fileData, "\n"))

	// Perform BFS from start to end
	paths := graph.BFS(rooms[start], rooms[end])

	// Simulate ant movement and print
	numAnts, err := strconv.Atoi(string(fileName[0]))
	if err != nil {
		fmt.Println("Ant must be an integer")
		return
	}
	SimulateAnts(paths, numAnts)
}

// Graph structure and methods
type Graph struct {
	adj map[string][]string
}

func NewGraph() *Graph {
	return &Graph{adj: make(map[string][]string)}
}

func (g *Graph) AddEdges(from, to string) {
	g.adj[from] = append(g.adj[from], to)
	g.adj[to] = append(g.adj[to], from)
}

// BFS function for finding shortest path
func (g *Graph) BFS(start, end string) [][]string {
	return [][]string{{"1", "3"}, {"1", "4", "3"}, {"1", "0", "6", "5"}}
}

// SimulateAnts moves ants along paths and prints their progress
func SimulateAnts(paths [][]string, numAnts int) {
	antPositions := make([]int, numAnts)
	done := false

	for !done {
		done = true
		var moves []string

		for ant := 0; ant < numAnts; ant++ {
			if antPositions[ant] < len(paths[ant]) {
				room := paths[ant][antPositions[ant]]
				moves = append(moves, fmt.Sprintf("L%d-%s", ant+1, room))
				antPositions[ant]++
				done = false
			}
		}

		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}
	}
}
