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
	numAnts, err := strconv.Atoi(strings.TrimSpace(string(ants)))
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
	var paths [][]string

	// A queue to keep track of the paths to explore
	queue := [][]string{{start}}
	visited := map[string]bool{start: true}

	for len(queue) > 0 {
		// Dequeue the first path from the queue
		path := queue[0]
		queue = queue[1:]

		// Get the last node in the current path
		node := path[len(path)-1]

		// If the last node is the end node, save the path
		if node == end {
			paths = append(paths, path)
			continue
		}

		// Explore the neighbors of the current node
		for _, neighbor := range g.adj[node] {
			// Avoid cycles by ensuring the neighbor is not already in the path
			if !visited[neighbor] {
				newPath := append([]string{}, path...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
				visited[neighbor] = true
			}
		}
	}

	return paths
}

func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// SimulateAnts moves ants along paths and prints their progress
func SimulateAnts(paths [][]string, numAnts int) {
    if len(paths) == 0 {
        fmt.Println("Error: No valid paths available.")
        return
    }

    // Each ant is assigned a path
    antPaths := make([][]string, numAnts)
    for i := 0; i < numAnts; i++ {
        antPaths[i] = paths[i%len(paths)]
    }

    // Initialize positions for each ant (-1 means not started)
    antPositions := make([]int, numAnts)
    for i := range antPositions {
        antPositions[i] = -1
    }

    // Track active rooms
    roomOccupancy := make(map[string]bool)

    finishedAnts := 0

    // Simulate the movement
    for finishedAnts < numAnts {
        var moves []string

        for ant := 0; ant < numAnts; ant++ {
            if antPositions[ant] == len(antPaths[ant])-1 { // Ant has finished its path
                continue
            }

            nextPosition := antPositions[ant] + 1
            nextRoom := antPaths[ant][nextPosition]

            // Check if the next room is available
            if !roomOccupancy[nextRoom] {
                // Move the ant to the next room
                if antPositions[ant] >= 0 { // Free the current room
                    currentRoom := antPaths[ant][antPositions[ant]]
                    delete(roomOccupancy, currentRoom)
                }

                moves = append(moves, fmt.Sprintf("L%d-%s", ant+1, nextRoom))
                roomOccupancy[nextRoom] = true
                antPositions[ant]++

                // Check if this ant has now finished
                if antPositions[ant] == len(antPaths[ant])-1 {
                    finishedAnts++
                }
            }
        }

        // Print the moves for this step
        if len(moves) > 0 {
            fmt.Println(strings.Join(moves, " "))
        }
    }
}


