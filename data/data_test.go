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
