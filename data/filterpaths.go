package data

func FilterPath(paths [][]string, farm *AntFarm) [][]string {
	var usedPaths [][]string
	// exlude the start room
	for i := 0; i < len(paths); i++ {
		paths[i] = paths[i][1:]
	}

	// place 'ants' on the rooms based on the size of the path
	track := farm.Ants
	for track > 0 {
		min := len(paths[0])
		index := 0
		for i := 1; i < len(paths); i++ {
			if len(paths[i]) < min {
				index = i
			}
		}
		paths[index] = append(paths[index], "ant")
		track--
	}

	//useable paths for the ants
	for i := 0; i < len(paths); i++ {
		if paths[i][len(paths[i])-1] == "ant" {
			usedPaths = append(usedPaths, paths[i])
		}
	}

	//calculate the turns
	var max = 0
	for i := 0; i < len(usedPaths); i++ {
		if len(usedPaths[i]) > max {
			max = len(usedPaths[i])
		}
	}

	farm.Turns = max - 1

	// remove the 'ants' from the useable paths
	for i := 0; i < len(usedPaths); i++ {
		for j := 0; j < len(usedPaths[i]); j++ {
			if usedPaths[i][j] == "ant" {
				usedPaths[i] = usedPaths[i][:j]
				break
			}
		}
	}
	return usedPaths
}
