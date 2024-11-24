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
