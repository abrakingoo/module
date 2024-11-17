package data

import(
	"lemin/utils"
)
// Graph structure with Nodes
type Graph struct {
	Node map[string][]string
}

// AddEdges adds a bidirectional edge between two nodes in the graph.
func (g *Graph) AddEdges(from, to string) {
	g.Node[from] = append(g.Node[from], to)
	g.Node[to] = append(g.Node[to], from)
}

// NewGraph initializes a new Graph and returns a pointer to it
func NewGraph() *Graph {
	return &Graph{Node: map[string][]string{}}
}


func (g *Graph) BFS(start, end string) [][]string {
	allPaths := [][]string{}
	qeue := [][]string{{start}}

	for len(qeue) > 0 {
		path := qeue[0]
		qeue = qeue[1:]
		node := path[len(path)-1]
	
		if node == end {
			allPaths = append(allPaths, path)
		} else {
			for _, neighbour := range g.Node[node] {
				if !utils.Contains(path, neighbour) {
					newPath := append([]string{}, path...)
					newPath = append(newPath, neighbour)
					qeue = append(qeue, newPath)
				}
			}
		}

	}

	return findMaxCompatiblePaths(allPaths)
}

// Simple helper to check if a path is compatible with a set of paths
func isCompatibleWithAll(path []string, pathSet [][]string) bool {
    // For each path in our current set
    for _, existingPath := range pathSet {
        // Check middle nodes (exclude start/end)
        for i := 1; i < len(path)-1; i++ {
            for j := 1; j < len(existingPath)-1; j++ {
                if path[i] == existingPath[j] {
                    return false
                }
            }
        }
    }
    return true
}

// Simplified version to find maximum compatible paths
func findMaxCompatiblePaths(paths [][]string) [][]string {
    maxSet := [][]string{}
    
    // Try each path as a starting point
    for i := 0; i < len(paths); i++ {
        currentSet := [][]string{paths[i]}
        
        // Try to add each remaining path
        for j := 0; j < len(paths); j++ {
            if i == j {
                continue
            }
            
            // If this path is compatible with all current paths, add it
            if isCompatibleWithAll(paths[j], currentSet) {
                currentSet = append(currentSet, paths[j])
            }
        }
        
        // Update maxSet if we found a larger set
        if len(currentSet) > len(maxSet) {
            maxSet = make([][]string, len(currentSet))
            copy(maxSet, currentSet)
        }
    }
    
    return maxSet
}
