package utils

func IsUnique(test []string, allArr [][]string) bool {
	for _, existingPath := range allArr {
		// Check middle nodes (exclude start/end)
		for i := 1; i < len(test)-1; i++ {
			for j := 1; j < len(existingPath)-1; j++ {
				if test[i] == existingPath[j] {
					return false
				}
			}
		}
	}
	return true

}
