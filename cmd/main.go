package main

import (
	"fmt"
	"log"
	//"log"
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
	ants := fileData[0]
	if len(ants) == 1 {
		fileData = fileData[1:]
	} else {
		log.Fatalf("ERROR: invalid data format")
	}

	// if len(fileData) == 0 || fileData[0] != "##start" {
	// 	log.Fatal("ERROR: invalid data format, no start room found")
	// }

	var (
		start int
		end int
		initial int
		count int
		rooms []string
		tunnels []string
	)

	// Parse rooms and find the start and end positions
	for i := 0; i < len(fileData); i++ {
		line := fileData[i]
		if line == "" {
			continue
		}

		if string(line[0]) != "#" {
			count++
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
					break
				}
				if string(fileData[i][0]) != "#" {
					rooms = append(rooms, string(fileData[i][0]))
				}
				i++
			}
			break
		}
		rooms = append(rooms, string(line[0])) // Add the room name
	}

	fmt.Println("the number of rooms ", rooms)
	// Parse tunnels starting from the "start" index
	for i := initial; i < len(fileData); i++ {
		line := fileData[i]
		if line == "" || line[0] == '#' {
			continue
		}
		tunnels = append(tunnels, line)
	}

	graph := data.NewGraph()
	fmt.Println("the tunnels are this ", tunnels)
	for _, connection := range tunnels {
		if strings.Contains(connection, "-") {
			rooms := strings.Split(connection, "-")
			graph.AddEdges(rooms[0], rooms[1])
		}
	}
	paths := graph.BFS(rooms[start],rooms[end])
	fmt.Println(paths)
}
