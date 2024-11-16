package main

import (
	"fmt"
	"log"
	"lemin/data"
	"lemin/utils"
	"os"
	"strings"
	"strconv"
)

type ant struct {
	index int
	target int
	targetIndex int
	iteration int
}

func newAnt() *ant {
	return &ant{}
}

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
	// if len(ants) == 1 && len(fileData[1]) != 1 {
	fileData = fileData[1:]
	// } else {
	// 	log.Fatalf("ERROR: invalid data format")
	// }
	antNum, err := strconv.Atoi(ants)
	if err != nil {
		log.Fatal("ERROR: invalid ant number format")
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
					r := strings.Split(fileData[i], " ")
					rooms = append(rooms, r[0])
				}
				i++
			}
		}
		r := strings.Split(line, " ")
		rooms = append(rooms, r[0]) // Add the room name
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
	
	paths := graph.BFS(rooms[start],rooms[end])
	
	for i:= 0 ; i< len(paths);i++ {
		paths[i] = paths[i][1:]
	}
	
	copy_path := make([][]string , len(paths))
	_= copy(copy_path,paths)


	track := antNum
	for antNum > 0 {
		min := len(copy_path[0])
		index := 0
		for i:= 1; i < len(copy_path); i++ {
			if len(copy_path[i]) < min {
				index = i
			}
		}
		copy_path[index] = append(copy_path[index], "ant")
		antNum--
	}

	var usedPaths [][]string 

	//useable paths for the ants
	for i:= 0; i < len(copy_path) ;i++ {
		if copy_path[i][len(copy_path[i])-1] == "ant" {
			usedPaths = append(usedPaths, copy_path[i])
		}
	}



	antmove := make([]*ant, antNum)

	for i:= 0; i < track; i++ {
		a := newAnt()
		a.index = i + 1
		antmove = append(antmove, a)
	}

	var iteration = 0
	var max = 0
	for i:= 0; i < len(usedPaths);i++ {
		if len(usedPaths[i]) > max {
			max = len(usedPaths[i])
		}
	}

	iteration = max - 1

	trackRoom := make(map[int]int)

	
	for i:= 0; i < len(usedPaths); i++ {
		for j:= 0; j < len(usedPaths[i]); j++ {
			if usedPaths[i][j] == "ant" {
				usedPaths[i] = usedPaths[i][:j]
				break
			}
		}
	} 

	for i:= 0; i < len(usedPaths); i++ {
		trackRoom[i] = (iteration + 1) - len(usedPaths[i])
	}



	var index = 0
	for i:= 0;  i < iteration; i++ {
		for j := 0; j < len(usedPaths); j++ {
			if index >= len(antmove) {
				break
			}
			if trackRoom[j] == 0 {
				continue
			}
			antmove[index].target = j
			antmove[index].iteration = i
			antmove[index].targetIndex = 0
			trackRoom[j]--
			index++
		}
	}


	for i:= 0; i < iteration; i++ {
		for j:= 0;  j < len(antmove); j++ {
			if antmove[j].iteration == i  && antmove[j].targetIndex < len(usedPaths[antmove[j].target])  {
				fmt.Printf("L%d-%s ", antmove[j].index, usedPaths[antmove[j].target][antmove[j].targetIndex])
				antmove[j].iteration += 1
				antmove[j].targetIndex += 1
			}
		}
		fmt.Println()
	}

}
