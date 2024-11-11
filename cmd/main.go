package main

import (
	"fmt"
	"log"
	"os"
	"strings"


	"lemin/utils"
	"lemin/data"
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
	utils.CheckError(err)

	fileData := strings.Split(string(file), "\n")
	if len(fileData) == 0 || fileData[0] != "##start" {
		log.Fatal("ERROR: invalid data format, no start room found")
	}

	var rooms []string
	var tunnels []string
	start := 0

	// Parse rooms and find the start and end positions
	for i := 1; i < len(fileData); i++ {
		line := fileData[i]
		if line == "" {
			continue
		}
		if line[0] == '#' && line != "##end" {
			continue
		}
		if line == "##end" {
			if i+1 < len(fileData) {
				start = i + 2
				rooms = append(rooms, string(fileData[i+1][0])) // Add the end room
			}
			break
		}
		rooms = append(rooms, string(line[0])) // Add the room name
	}

	fmt.Println(rooms)
	// Parse tunnels starting from the "start" index
	for i := start; i < len(fileData); i++ {
		line := fileData[i]
		if line == "" || line[0] == '#' {
			continue
		}
		tunnels = append(tunnels, line)
	}

	graph := data.NewGraph()

	for _, connection := range tunnels {
		rooms := strings.Split(connection, "-")
		graph.AddEdges(rooms[0], rooms[1])
	}

	paths := graph.BFS("1", "0")

	fmt.Println(paths)

}
