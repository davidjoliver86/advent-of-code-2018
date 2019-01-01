package main

import (
	"fmt"
)

// Initial input was "030121"
// The first star expects us to use "030121" as an integer - that is: "30,121 iterations"
// The second star expects a sequence of {0, 3, 0, 1, 2, 1}
// Also the initial two recipes (3, 7) are considered constant
const inputIterations = 30121

var initialRecipes = []int{3, 7}
var inputSequence = []int{0, 3, 0, 1, 2, 1}

// Elf is a container with an index that points to the recipe array
type Elf struct {
	// Current is the index of the recipe array
	id      int
	Current int
}

func (e *Elf) String() string {
	return fmt.Sprintf("Elf[%d]=>(%d)", e.id, e.Current)
}

// Score deferences the recipe array at the Elf's Current location
func (e *Elf) Score(r *[]int) int {
	score := (*r)[e.Current]
	return score
}

// AddRecipe combines the Scores of the teo Elfs' recipes and adds each digit to the array of recipe scores
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

// FindSequence continually adds recipes until seq appears at the tail-end of the recipe array, then returns
// the number of recipes before that (in other words, the index *of* seq's first appearance)
func FindSequence(recipes *[]int, e1 *Elf, e2 *Elf, seq []int) int {
	counter := 0
	index := 0
	sLen := len(seq)
	for {
		AddRecipe(recipes, e1, e2)
		// compare consecutive digits from recipes[index] onward
		// as long as digits match, increment the couter
		// once they dont match, reset counter to 0
		for i, v := range seq {
			if (*recipes)[index+i] == v {
				counter++
				if counter == sLen {
					return index
				}
			} else {
				index++
				counter = 0
				break
			}
		}
	}
}

func firstStar() {
	// first star
	recipes := initialRecipes[:]
	e1 := Elf{0, 0}
	e2 := Elf{1, 1}
	for len(recipes) < (inputIterations + 10) {
		AddRecipe(&recipes, &e1, &e2)
	}
	fmt.Println(recipes[inputIterations:])
}

func secondStar() {
	// secondstar
	recipes := initialRecipes[:]
	e1 := Elf{0, 0}
	e2 := Elf{1, 1}
	seq := FindSequence(&recipes, &e1, &e2, inputSequence)
	fmt.Println(seq)
}

func main() {
	firstStar()
	secondStar()
}
