package main

import (
	"fmt"
	"testing"
)

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
