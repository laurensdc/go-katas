package life

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_live_cell_with_fewer_than_two_live_neighbours_dies(t *testing.T) {
	before := [][]bool{
		{false, false, true, false},
		{false, true, false, true},
	}

	after := [][]bool{
		{false, false, true, false},
		{false, false, true, false},
	}

	result := Tick(before)

	fmt.Printf("%v\n", result)

	if !reflect.DeepEqual(result, after) {
		t.Errorf("Expected\n%v\nGot\n%v\n", after, result)
	}
}

func Test_live_cell_with_more_than_three_live_neighbours_dies(t *testing.T) {
	before := [][]bool{
		{true, true, true, false},
		{false, true, true, true},
	}

	after := [][]bool{
		{true, false, false, true},
		{true, false, false, true},
	}

	result := Tick(before)

	if !reflect.DeepEqual(result, after) {
		t.Errorf("Expected\n%v\nGot\n%v\n", after, result)
	}
}

func Test_live_cell_with_two_or_three_live_neighbours_lives_on(t *testing.T) {
	before := [][]bool{
		{false, true, true, false},
		{true, false, true, false},
		{true, true, true, true},
	}

	after := [][]bool{
		{false, true, true, false},
		{true, false, false, false},
		{true, false, true, true},
	}

	result := Tick(before)

	if !reflect.DeepEqual(result, after) {
		t.Errorf("Expected\n%v\nGot\n%v\n", after, result)
	}
}

func Test_any_dead_cell_with_exactly_three_live_neighbours_becomes_a_live_cell(t *testing.T) {
	before := [][]bool{
		{false, false, true, true},
		{true, true, false, false},
		{false, false, false, true},
	}

	after := [][]bool{
		{false, true, true, false},
		{false, true, false, true},
		{false, false, false, false},
	}

	result := Tick(before)

	if !reflect.DeepEqual(result, after) {
		t.Errorf("Expected\n%v\nGot\n%v\n", after, result)
	}

}

func Test_count_neighbors(t *testing.T) {
	state := [][]bool{
		{false, false, false, true},
		{true, false, true, true},
		{false, true, false, true},
	}

	if val := countNeighbors(state, 0, 0); val != 1 {
		t.Errorf("Expected 0, 0 to have 1 neighbors, got %v\n", val)
	}

	if val := countNeighbors(state, 1, 3); val != 3 {
		t.Errorf("Expected 1, 3 to have 3 neighbors, got %v\n", val)
	}

	if val := countNeighbors(state, 1, 2); val != 4 {
		t.Errorf("Expected 1, 2 to have 4 neighbors, got %v\n", val)
	}

	if val := countNeighbors(state, 1, 0); val != 1 {
		t.Errorf("Expected 1, 0 to have 1 neighbor, got %v\n", val)
	}
}
