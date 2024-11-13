package main

import (
	"fmt"
	"log"
	"lemin/data"
	"lemin/utils"
	"os"
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
	utils.CheckError(err)

	fileData := strings.Split(string(file), "\n")
	//check if the number of ants is provided
	ants := fileData[0]
	if len(ants) == 1 && len(fileData[1]) != 1 {
		fileData = fileData[1:]
	} else {
		log.Fatalf("ERROR: invalid data format")
	}

	//check if the progam has start or end
	var first, last bool 

	for i:=0 ; i < len(fileData); i++ {
		if string(fileData[i]) == "##start" {
			first = true
		}

		if string(fileData[i]) == "##end" {
			last = true
			if first {
				break// if both are true
			}
		}
	}

	if !first || !last {
		log.Fatalf("Error: invalid data format")
	}

	var (
		start int
		end int
		initial int
		count int
		rooms []string
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

	graph := data.NewGraph()
	for _, connection := range tunnels {
		if strings.Contains(connection, "-") {
			rooms := strings.Split(connection, "-")
			graph.AddEdges(rooms[0], rooms[1])
		}
	}
	fmt.Println("this is the graph ", graph)
	paths := graph.BFS(rooms[start],rooms[end])
	fmt.Println(paths)
}
