package main

import (
	"fmt"
	"lemin/data"
	"lemin/utils"
	"log"
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
	if !data.CheckValidFormat(farm) {
		log.Fatal("Error: invalid file format, missing or wrong values")
	}

	graph := data.NewGraph()
	if len(farm.Tunnels) == 0 {
		log.Fatal("ERROR: no valid tunnels, check file format")
	}
	for _, connection := range farm.Tunnels {
		if strings.Contains(connection, "-") {
			rooms := strings.Split(connection, "-")
			graph.AddEdges(rooms[0], rooms[1])
		}
	}

	paths := graph.BFS(farm.Start, farm.End)
	if len(paths) == 0 {
		log.Fatal("ERROR: no valid paths present in the file provided")
	}
	usedPaths := data.FilterPath(paths, farm)
	data.PrintResult(usedPaths, farm)

}
