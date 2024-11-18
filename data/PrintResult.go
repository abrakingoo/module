package data

import (
	"fmt"
	"strings"
)

func PrintResult(paths [][]string, farm *AntFarm) {

	trackRoom := make(map[int]int)
	// track room capacity
	for i := 0; i < len(paths); i++ {
		trackRoom[i] = (farm.Turns + 1) - len(paths[i])
	}

	track := farm.Ants
	antArray := make([]*Ant, farm.Ants)

	for i := 0; i < track; i++ {
		a := NewAnt()
		a.Index = i + 1
		antArray[i] = a
	}

	// assign which turn the ant will start to be printed
	// assign which path to target
	// target index will always start at 0
	var index = 0
	for i := 0; i < farm.Turns; i++ {
		for j := 0; j < len(paths); j++ {
			if index >= len(antArray) {
				break
			}
			if trackRoom[j] == 0 {
				continue
			}
			antArray[index].Target = j
			antArray[index].Iteration = i
			antArray[index].TargetIndex = 0
			trackRoom[j]--
			index++
		}
	}

	var result string
	for i := 0; i < farm.Turns; i++ {
		for j := 0; j < len(antArray); j++ {
			// check if the iteration is == to turns, and it has not exceeded the target paths index
			if antArray[j].Iteration == i && antArray[j].TargetIndex < len(paths[antArray[j].Target]) {
				result += fmt.Sprintf("L%d-%s ", antArray[j].Index, paths[antArray[j].Target][antArray[j].TargetIndex])
				//always add the iteration and target index so the ants will be printed till they reach the last room
				antArray[j].Iteration += 1
				antArray[j].TargetIndex += 1
			}
		}
		result = strings.TrimRight(result, " ")
		result += "\n"
	}
	fmt.Println(strings.Join(farm.Filedata, "\n"))
	fmt.Println()
	fmt.Print(result)
}
