package main

import "testing"

type testcaseArea struct {
	points   []Point
	target   Point
	expected int
}

var example = testcaseArea{
	points: []Point{
		{1, 1}, // A
		{1, 6}, // B
		{8, 3}, // C
		{3, 4}, // D
		{5, 5}, // E
		{8, 9}, // F
	},
	expected: 17,
}

func TestArea(t *testing.T) {
	grid := Grid(example.points, 10)
	areas := Area(grid)
	if areas[4] != example.expected {
		t.Errorf("Expected area of point E to be %d, got %d", example.expected, areas[4])
	}
}
