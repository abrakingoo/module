package utils

func CheckStartAndEnd(file []string) bool {
	var (
		first = false
		last = false
	)

	for i := 0; i < len(file); i++ {
		if string(file[i]) == "##start" {
			if first  {
				return false
			} else {
				first = true
			}
		}

		if string(file[i]) == "##end" {
			if last {
				return false
			} else {
				last = true
			}		
		}
	}

	if !first || !last {
		return false
	}
	return true
}
