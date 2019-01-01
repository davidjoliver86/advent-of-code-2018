package main

import (
	"fmt"
	"testing"
)

type findSequenceTestCase struct {
	input    []int
	expected int
}

var findSequenceTestCases = []findSequenceTestCase{
	{[]int{5, 1, 5, 8, 9}, 9},
	{[]int{0, 1, 2, 4, 5}, 5},
	{[]int{9, 2, 5, 1, 0}, 18},
	{[]int{5, 9, 4, 1, 4}, 2018},
}

func TestAddRecipe(t *testing.T) {
	recipes := []int{3, 7}
	e1 := Elf{0, 0}
	e2 := Elf{1, 1}
	for len(recipes) < 19 {
		AddRecipe(&recipes, &e1, &e2)
	}
	result := fmt.Sprintf("%v", recipes[9:])
	if result != "[5 1 5 8 9 1 6 7 7 9]" {
		t.Errorf("Expected [5 1 5 8 9 1 6 7 7 9], got %v", result)
	}
}

func TestFindSequence(t *testing.T) {
	for _, testcase := range findSequenceTestCases {
		recipes := []int{3, 7}
		e1 := Elf{0, 0}
		e2 := Elf{1, 1}
		seq := FindSequence(&recipes, &e1, &e2, testcase.input)
		fmt.Println(seq)
		if seq != testcase.expected {
			t.Errorf("Expected to find %v at %d, got %d", testcase.input, testcase.expected, seq)
		}
	}
}
