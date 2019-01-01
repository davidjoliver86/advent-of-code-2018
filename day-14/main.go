package main

import "fmt"

const inputIterations = 30121

var initialRecipes = []int{3, 7}
var inputSequence = []int{0, 3, 0, 1, 2, 1}

type Elf struct {
	// Current is the index of the recipe array
	id      int
	Current int
}

func (e *Elf) String() string {
	return fmt.Sprintf("Elf[%d]=>(%d)", e.id, e.Current)
}

func (e *Elf) Score(r *[]int) int {
	score := (*r)[e.Current]
	return score
}

func AddRecipe(recipes *[]int, e1 *Elf, e2 *Elf) {
	sum := e1.Score(recipes) + e2.Score(recipes)
	if sum >= 10 {
		*recipes = append(*recipes, sum/10)
	}
	*recipes = append(*recipes, sum%10)

	// how many spaces does each elf move? (1 + score of current)
	for _, elf := range []*Elf{e1, e2} {
		spaces := 1 + elf.Score(recipes)
		new := (elf.Current + spaces) % len(*recipes)
		elf.Current = new
	}
}

func main() {
	// first star
	recipes := initialRecipes[:]
	e1 := Elf{0, 0}
	e2 := Elf{1, 1}
	for len(recipes) < (inputIterations + 10) {
		AddRecipe(&recipes, &e1, &e2)
	}
	fmt.Println(recipes[inputIterations:])

	// secondstar
	recipes := initialRecipes[:]
	e1 := Elf{0, 0}
	e2 := Elf{1, 1}
}
