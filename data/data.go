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
	
		if node == end && utils.IsUnique(allPaths, path[1:len(path)-1]){
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

	return allPaths

}