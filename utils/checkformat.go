package utils


func CheckStartAndEnd(file []string) bool {
	var first, last bool 

	for i:=0 ; i < len(file); i++ {
		if string(file[i]) == "##start" {
			first = true
		}

		if string(file[i]) == "##end" {
			last = true
			if first {
				break// if both are true
			}
		}
	}

	if !first || !last {
		return false
	}
	return true
}