package data

import (
	"log"
	"strconv"
	"strings"

	"lemin/utils"
)

func GetFileData(file []byte, farm *AntFarm) {
	var (
		start   int
		end     int
		initial int
		count   int
		rooms   []string
		tunnels []string
	)

	fileData := strings.Split(string(file), "\n")
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

loop:
	for i := 1; i < len(fileData); i++ {
		line := fileData[i]
		if line == "" {
			continue
		}
		if line == "##start" {
			start = count // index of the start room
		}

		if line[0] == '#' && line != "##end" {
			continue
		}

		if line == "##end" {
			end = count // index of the end room
			for i < len(fileData) {
				if strings.Contains(string(fileData[i]), "-") {
					initial = i // track to where the tunnels start
					break loop
				}
				if string(fileData[i][0]) != "#" {
					rooms = append(rooms, fileData[i])
				}
				i++
			}
		}
		rooms = append(rooms, line) // Add the room name
		count++
	}

	for i := initial; i < len(fileData); i++ {
		line := fileData[i]
		if line == "" || line[0] == '#' {
			continue
		}
		tunnels = append(tunnels, line)
	}

	farm.Ants = ants
	farm.Filedata = fileData
	farm.Start = start
	farm.End = end
	farm.Rooms = rooms
	farm.Tunnels = tunnels
}
