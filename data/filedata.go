package data

import (
	"log"
	"strconv"
	"strings"

	"lemin/utils"
)

func GetFileData(file []byte, farm *AntFarm) {
	fileData := strings.Split(strings.TrimSpace(string(file)), "\n")
	farm.Filedata = fileData

	ants, err := strconv.Atoi(fileData[0])
	utils.CheckError(err)
	// if number of ants is 0
	if ants <= 0 {
		log.Fatal("ERROR: invalid number of ants")
	}
	farm.Ants = ants
	// exclude number of ants
	
	if !utils.CheckStartAndEnd(fileData[1:]) {
		log.Fatal("ERROR: invalid file format, missing or repeated ##start and ##end values in the present file")
	}

	for i := 1; i < len(fileData); i++ {
		line := fileData[i]
		if line == ""  {
			continue
		}
		if strings.Contains(line, "#") {
			if line == "##start" {
				farm.Start = strings.Split(fileData[i+1], " ")[0]
				continue
			} else if line == "##end" {
				farm.End = strings.Split(fileData[i+1], " ")[0]
				continue
			} else {
				continue
			}
		}

		if strings.Contains(line, "-") {
			farm.Tunnels = append(farm.Tunnels, line)
			continue
		}
		farm.Rooms = append(farm.Rooms, line) // Add the room name
	}

}
