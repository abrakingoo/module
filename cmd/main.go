package main

import (
	"fmt"
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

	if !strings.HasSuffix(fileName, ".txt") {
		fmt.Println("ERROR: invalid file format, only (.txt) files are supported.")
		return
	}

	// Read the contents of the file
	file, err := os.ReadFile(fileName)
	utils.CheckError(err)
	farm := data.NewFarm()

	data.GetFileData(file, farm)

	graph := data.NewGraph()
	for _, connection := range farm.Tunnels {
		if strings.Contains(connection, "-") {
			rooms := strings.Split(connection, "-")
			graph.AddEdges(rooms[0], rooms[1])
		}
	}

	paths := graph.BFS(farm.Rooms[farm.Start],farm.Rooms[farm.End])
	usedPaths := data.FilterPath(paths, farm) 
	data.PrintResult(usedPaths, farm)
		
	
}
