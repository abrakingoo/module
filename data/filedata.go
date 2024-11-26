package data

import (
	"log"
	"strconv"
	"strings"

	"lemin/utils"
)

func GetFileData(file []byte, farm *AntFarm) {
	var (
		start   string
		end     string
		rooms   []string
		tunnels []string
	)
	
	fileData := strings.Split(strings.TrimSpace(string(file)), "\n")
	ants, err := strconv.Atoi(fileData[0])
	utils.CheckError(err)
	// if number of ants is 0
	if ants <= 0 {
		log.Fatal("ERROR: invalid number of ants")
	}
	// exclude number of ants
	if !utils.CheckStartAndEnd(fileData[1:]) {
		log.Fatal("ERROR: invalid file format, missing ##start and ##end in file")
	}

	for i := 1; i < len(fileData); i++ {
		line := fileData[i]
		if line == "" || strings.Contains(line , "#") {
			continue
		}
		if line == "##start" {
			start = "" // index of the start room
		}

		if line[0] == '#' && line != "##end" {
			continue
		}

		if strings.Contains(line, "-") {
			rooms = append(rooms, line)
		}

		if line == "##end" {
			end = ""
			 // index of the end room
		}
		rooms = append(rooms, line) // Add the room name
	}

	

	farm.Ants = ants
	farm.Filedata = fileData
	farm.Start = start
	farm.End = end
	farm.Rooms = rooms
	farm.Tunnels = tunnels
}
