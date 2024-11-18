package data

import (
	"lemin/utils"
	"strconv"
	"strings"
	"log"
)


func GetFileData(file []byte , farm *AntFarm) {
	var (
		start int
		end int
		initial int
		count int
		rooms []string
		tunnels []string
	)

	fileData := strings.Split(string(file), "\n")
	ants, err := strconv.Atoi(fileData[0])
	utils.CheckError(err)
	
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
	
	for i := initial; i < len(fileData); i++ {
		line := fileData[i]
		if line == "" || line[0] == '#' {
			continue
		}
		tunnels = append(tunnels, line)
	}


	farm.Ants = ants
	farm.Farm = fileData
	farm.Start = start
	farm.End = end
	farm.Rooms = rooms
	farm.Tunnels = tunnels

}