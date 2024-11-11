package utils

func IsUnique(allArr [][]string, test []string) bool {
	for _, arr := range allArr {
		for _, val := range arr {
			for _, n := range test {
				if val == n {
					return false
				}
			}
		}
	}
	return true
}