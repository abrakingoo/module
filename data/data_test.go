package data

import (
	"testing"
)

// TestNewFarm verifies that NewFarm constructor creates an AntFarm with zero values
func TestNewFarm(t *testing.T) {
	farm := NewFarm()
	
	// Check if farm is created
	if farm == nil {
		t.Fatal("NewFarm() returned nil")
	}

	// Verify all fields are initialized to zero values
	if farm.Ants != 0 {
		t.Errorf("NewFarm().Ants = %d, want 0", farm.Ants)
	}
	if farm.Start != 0 {
		t.Errorf("NewFarm().Start = %d, want 0", farm.Start)
	}
	if farm.End != 0 {
		t.Errorf("NewFarm().End = %d, want 0", farm.End)
	}
	if farm.Turns != 0 {
		t.Errorf("NewFarm().Turns = %d, want 0", farm.Turns)
	}
	
	// Verify slices are initialized to empty slices
	if farm.Filedata == nil {
		t.Error("NewFarm().Filedata = nil, want empty slice")
	}
	if farm.Rooms == nil {
		t.Error("NewFarm().Rooms = nil, want empty slice")
	}
	if farm.Tunnels == nil {
		t.Error("NewFarm().Tunnels = nil, want empty slice")
	}
	if farm.Paths == nil {
		t.Error("NewFarm().Paths = nil, want empty slice")
	}
}

func TestNewFarmBasic(t *testing.T) {
	farm := NewFarm()
	if farm == nil {
		t.Error("Expected NewFarm() to return a non-nil pointer to AntFarm")
	}
	
	// Test initial values
	if farm.Ants != 0 {
		t.Errorf("Expected initial Ants to be 0, got %d", farm.Ants)
	}
	if farm.Start != 0 {
		t.Errorf("Expected initial Start to be 0, got %d", farm.Start)
	}
	if farm.End != 0 {
		t.Errorf("Expected initial End to be 0, got %d", farm.End)
	}
}

func TestNewAnt(t *testing.T) {
	ant := NewAnt()
	if ant == nil {
		t.Error("Expected NewAnt() to return a non-nil pointer to Ant")
	}
	
	// Test initial values
	if ant.Index != 0 {
		t.Errorf("Expected initial Index to be 0, got %d", ant.Index)
	}
	if ant.Target != 0 {
		t.Errorf("Expected initial Target to be 0, got %d", ant.Target)
	}
	if ant.TargetIndex != 0 {
		t.Errorf("Expected initial TargetIndex to be 0, got %d", ant.TargetIndex)
	}
	if ant.Iteration != 0 {
		t.Errorf("Expected initial Iteration to be 0, got %d", ant.Iteration)
	}
}

// Test the Ant struct functionality
func TestAntFunctionality(t *testing.T) {
	ant := NewAnt()
	
	// Test setting values
	ant.Index = 1
	ant.Target = 2
	ant.TargetIndex = 3
	ant.Iteration = 4
	
	if ant.Index != 1 {
		t.Errorf("Ant.Index = %v, want 1", ant.Index)
	}
	if ant.Target != 2 {
		t.Errorf("Ant.Target = %v, want 2", ant.Target)
	}
	if ant.TargetIndex != 3 {
		t.Errorf("Ant.TargetIndex = %v, want 3", ant.TargetIndex)
	}
	if ant.Iteration != 4 {
		t.Errorf("Ant.Iteration = %v, want 4", ant.Iteration)
	}
}

// Test the AntFarm struct functionality
func TestAntFarmFunctionality(t *testing.T) {
	farm := NewFarm()
	
	// Test setting values
	farm.Ants = 10
	farm.Start = 1
	farm.End = 2
	farm.Turns = 5
	farm.Filedata = []string{"test1", "test2"}
	farm.Rooms = []string{"room1", "room2"}
	farm.Tunnels = []string{"tunnel1", "tunnel2"}
	farm.Paths = []string{"path1", "path2"}
	
	if farm.Ants != 10 {
		t.Errorf("Farm.Ants = %v, want 10", farm.Ants)
	}
	if farm.Start != 1 {
		t.Errorf("Farm.Start = %v, want 1", farm.Start)
	}
	if farm.End != 2 {
		t.Errorf("Farm.End = %v, want 2", farm.End)
	}
	if farm.Turns != 5 {
		t.Errorf("Farm.Turns = %v, want 5", farm.Turns)
	}
	if len(farm.Filedata) != 2 {
		t.Errorf("len(Farm.Filedata) = %v, want 2", len(farm.Filedata))
	}
	if len(farm.Rooms) != 2 {
		t.Errorf("len(Farm.Rooms) = %v, want 2", len(farm.Rooms))
	}
	if len(farm.Tunnels) != 2 {
		t.Errorf("len(Farm.Tunnels) = %v, want 2", len(farm.Tunnels))
	}
	if len(farm.Paths) != 2 {
		t.Errorf("len(Farm.Paths) = %v, want 2", len(farm.Paths))
	}
}

func TestGetFileData(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantAnts int
		wantRooms int
		wantTunnels int
	}{
		{
			name: "Basic valid input",
			input: `5
##start
room0 0 0
##end
room1 1 1
room2 2 2
room0-room1
room1-room2`,
			wantAnts: 5,
			wantRooms: 3,
			wantTunnels: 2,
		},
		{
			name: "Input with comments",
			input: `10
#comment here
##start
room0 0 0
##end
room1 1 1
#comment
room2 2 2
room0-room1
#tunnel comment
room1-room2`,
			wantAnts: 10,
			wantRooms: 3,
			wantTunnels: 2,
		},
		{
			name: "Multiple tunnels",
			input: `3
##start
room0 0 0
##end
room1 1 1
room2 2 2
room3 3 3
room0-room1
room1-room2
room2-room3
room3-room1`,
			wantAnts: 3,
			wantRooms: 4,
			wantTunnels: 4,
		},
		{
			name: "Complex room names",
			input: `7
##start
start_room 0 0
##end
end_room 1 1
middle_room 2 2
another_room 3 3
start_room-middle_room
middle_room-end_room
another_room-end_room`,
			wantAnts: 7,
			wantRooms: 4,
			wantTunnels: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			farm := NewFarm()
			GetFileData([]byte(tt.input), farm)

			if farm.Ants != tt.wantAnts {
				t.Errorf("GetFileData() got Ants = %v, want %v", farm.Ants, tt.wantAnts)
			}
			if len(farm.Rooms) != tt.wantRooms {
				t.Errorf("GetFileData() got %v rooms, want %v", len(farm.Rooms), tt.wantRooms)
			}
			if len(farm.Tunnels) != tt.wantTunnels {
				t.Errorf("GetFileData() got %v tunnels, want %v", len(farm.Tunnels), tt.wantTunnels)
			}
			// Check if Start and End are set
			if farm.Start < 0 || farm.Start >= len(farm.Rooms) {
				t.Error("GetFileData() start room index out of bounds")
			}
			if farm.End < 0 || farm.End >= len(farm.Rooms) {
				t.Error("GetFileData() end room index out of bounds")
			}
			// Check if filedata is properly set
			if len(farm.Filedata) == 0 {
				t.Error("GetFileData() filedata not set properly")
			}
			// Verify that start and end are different
			if farm.Start == farm.End {
				t.Error("GetFileData() start and end rooms should be different")
			}
		})
	}
}

// TestNewGraph verifies that NewGraph creates a graph with initialized map
func TestNewGraph(t *testing.T) {
	g := NewGraph()
	
	if g == nil {
		t.Fatal("NewGraph() returned nil")
	}
	
	if g.Node == nil {
		t.Error("NewGraph().Node map not initialized")
	}
	
	if len(g.Node) != 0 {
		t.Errorf("NewGraph().Node should be empty, got %d items", len(g.Node))
	}
}

// TestGraphAddEdges verifies that edges are added correctly in both directions
func TestGraphAddEdges(t *testing.T) {
	g := NewGraph()
	
	// Test adding a single edge
	g.AddEdges("A", "B")
	
	// Check if edge A->B exists
	if len(g.Node["A"]) != 1 || g.Node["A"][0] != "B" {
		t.Error("Edge A->B not added correctly")
	}
	
	// Check if edge B->A exists (bidirectional)
	if len(g.Node["B"]) != 1 || g.Node["B"][0] != "A" {
		t.Error("Edge B->A not added correctly")
	}
	
	// Test adding multiple edges to same node
	g.AddEdges("A", "C")
	
	if len(g.Node["A"]) != 2 {
		t.Error("Second edge from A not added correctly")
	}
	
	// Verify all connections
	expectedA := []string{"B", "C"}
	for i, v := range expectedA {
		if g.Node["A"][i] != v {
			t.Errorf("Node A connections incorrect at index %d, got %s want %s", i, g.Node["A"][i], v)
		}
	}
}

// TestGraphBFS verifies breadth-first search path finding
func TestGraphBFS(t *testing.T) {
	tests := []struct {
		name          string
		edges         [][2]string
		start         string
		end           string
		wantPathCount int
	}{
		{
			name: "Simple path",
			edges: [][2]string{
				{"start", "A"},
				{"A", "end"},
			},
			start:         "start",
			end:           "end",
			wantPathCount: 1,
		},
		{
			name: "Multiple paths",
			edges: [][2]string{
				{"start", "A"},
				{"A", "B"},
				{"B", "end"},
				{"start", "C"},
				{"C", "end"},
			},
			start:         "start",
			end:           "end",
			wantPathCount: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGraph()
			
			// Build graph
			for _, edge := range tt.edges {
				g.AddEdges(edge[0], edge[1])
			}
			
			paths := g.BFS(tt.start, tt.end)
			
			if len(paths) != tt.wantPathCount {
				t.Errorf("BFS() found %d paths, want %d", len(paths), tt.wantPathCount)
			}
			
			// Verify each path starts and ends correctly
			for i, path := range paths {
				if len(path) < 2 {
					t.Errorf("Path %d is too short: %v", i, path)
					continue
				}
				if path[0] != tt.start {
					t.Errorf("Path %d doesn't start with %s: %v", i, tt.start, path)
				}
				if path[len(path)-1] != tt.end {
					t.Errorf("Path %d doesn't end with %s: %v", i, tt.end, path)
				}
			}
		})
	}
}
