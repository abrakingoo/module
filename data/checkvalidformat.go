package data

import (
	"strconv"
	"strings"
)

func CheckValidFormat(farm *AntFarm) bool {
	//check if start rooms and end rooms are the same
	
	if farm.Start == farm.End {
		return false
	}

	//check if co-ordinates are integers
	for i:= 0; i < len(farm.Rooms); i++ {
		data := strings.Split(strings.TrimSpace(farm.Rooms[i]), " ")
		//check if rooms start with an l
		if strings.ToLower(string(data[0][0])) == "l" {
			return false
		}
		// check if there more than two co-ordinates
		if len(data) != 3 { 
			return false
		}
		for j:= 1; j < len(data); j++ {
			if _, err := strconv.Atoi(data[j]); err != nil {
				return false
			}
		}
	}

	//check for repeated rooms
	var checker = make(map[string]int)
	for i:= 0; i < len(farm.Rooms); i++ {
		data := strings.Split(farm.Rooms[i], " ")[0]
		checker[data]++
	}

	for keys := range checker {
		if checker[keys] > 1 {
			return false
		}
	}


	 //check for tunnels if the start with an l
	for i:= 0; i < len(farm.Tunnels); i++ {
		data := strings.Split(farm.Tunnels[i], "-")
		//check if the tunnel format is wrong, like extra tunnels
		 if len(data) != 2 {
		 	return false
		 }
		for j :=0 ; j < len(data); j++ {
			if strings.ToLower(string(data[j][0])) == "#" || strings.ToLower(string(data[j][0])) == "l" {
				return false
			} 
		}
	}
	return true
}