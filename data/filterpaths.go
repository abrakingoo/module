package data

func FilterPath(paths [][]string, farm *AntFarm) [][]string {
	var usedPaths [][]string
	// exlude the start room
	for i := 0; i < len(paths); i++ {
		paths[i] = paths[i][1:]
	}

	copy_path := make([][]string, len(paths))
	_ = copy(copy_path, paths)

	// place 'ants' on the rooms based on the size of the path
	track := farm.Ants
	for track > 0 {
		min := len(copy_path[0])
		index := 0
		for i := 1; i < len(copy_path); i++ {
			if len(copy_path[i]) < min {
				index = i
			}
		}
		copy_path[index] = append(copy_path[index], "ant")
		track--
	}

	//useable paths for the ants
	for i := 0; i < len(copy_path); i++ {
		if copy_path[i][len(copy_path[i])-1] == "ant" {
			usedPaths = append(usedPaths, copy_path[i])
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
