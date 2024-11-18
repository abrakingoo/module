package utils

// Helper function to check if a node is already in the path
func Contains(path []string, node string) bool {
	for _, n := range path {
		if n == node {
			return true
		}
	}
	return false
}
