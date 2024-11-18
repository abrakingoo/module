package data 

import (
	"fmt"
	"strings"
)


func PrintResult(paths [][]string, farm *AntFarm) {
	trackRoom := make(map[int]int)
	for i:= 0; i < len(paths); i++ {
		trackRoom[i] = (farm.Turns + 1) - len(paths[i])
	}
	
	track := farm.Ants
	antArray := make([]*Ant, farm.Ants)
	
	for i:= 0; i < track; i++ {
		a := NewAnt()
		a.Index = i + 1
		antArray[i] = a
	}
	

	var index = 0
	for i:= 0;  i < farm.Turns; i++ {
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
	for i:= 0; i < farm.Turns; i++ {
		for j:= 0;  j < len(antArray); j++ {
			if antArray[j].Iteration == i  && antArray[j].TargetIndex < len(paths[antArray[j].Target])  {
				result += fmt.Sprintf("L%d-%s ", antArray[j].Index, paths[antArray[j].Target][antArray[j].TargetIndex])
				antArray[j].Iteration += 1
				antArray[j].TargetIndex += 1
			}
		}
		result = strings.TrimRight(result, " ")
		result += "\n"
	}
	fmt.Println(strings.Join(farm.Farm, "\n"))
	fmt.Println()
	fmt.Print(result)
}